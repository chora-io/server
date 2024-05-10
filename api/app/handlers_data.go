package app

import (
	"crypto"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/piprate/json-gold/ld"

	"github.com/chora-io/server/api/regen"
	db "github.com/chora-io/server/db/client"
)

func GetData(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	iri := vars["iri"]

	split := strings.Split(iri, ":")
	if len(split) == 1 {
		respondError(
			rw,
			http.StatusBadRequest,
			fmt.Sprintf("invalid iri: %s is not a valid iri", iri),
		)
		return
	}

	// other prefixes ok but stored with chora prefix on chora server
	iri = "chora:" + split[1]

	_, err := regen.ParseIRI(iri)
	if err != nil {
		respondError(
			rw,
			http.StatusBadRequest,
			fmt.Sprintf("invalid iri: %s is not a valid iri", iri),
		)
		return
	}

	d, err := dbr.GetData(r.Context(), iri)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(
				rw,
				http.StatusNotFound,
				fmt.Sprintf("data with iri %s does not exist", iri),
			)
		} else {
			respondError(rw, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondJSON(rw, http.StatusOK, NewGetDataResponse(iri, string(d.Jsonld)))
}

func PostData(dbw db.Writer, rw http.ResponseWriter, r *http.Request) {
	var req PostDataRequest

	bz, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(bz, &req)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	digest := strings.TrimSpace(req.Digest)
	if digest != "BLAKE2B_256" {
		respondError(rw, http.StatusBadRequest, "digest algorithm must be BLAKE2B_256")
		return
	}

	canon := strings.TrimSpace(req.Canon)
	if canon != ld.AlgorithmURDNA2015 {
		respondError(
			rw,
			http.StatusBadRequest,
			fmt.Sprintf("canonicalization algorithm must be %s", ld.AlgorithmURDNA2015),
		)
		return
	}

	merkle := strings.TrimSpace(req.Merkle)
	if merkle != "UNSPECIFIED" {
		respondError(rw, http.StatusBadRequest, "merkle tree must be UNSPECIFIED")
		return
	}

	processor := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/n-quads"
	options.Algorithm = ld.AlgorithmURDNA2015

	var doc interface{}

	err = json.Unmarshal([]byte(req.Jsonld), &doc)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	normalized, err := processor.Normalize(doc, options)
	if err != nil {
		respondError(rw, http.StatusBadRequest, err.Error())
		return
	}

	hash := crypto.BLAKE2b_256.New()
	_, err = hash.Write([]byte(normalized.(string)))
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	ch := regen.ContentHash_Graph{
		Hash:                      hash.Sum(nil),
		DigestAlgorithm:           regen.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: regen.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		MerkleTree:                regen.GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
	}

	// use chora prefix when stored on chora server
	iri, err := ch.ToIRI("chora")
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	jsonld := compactJSONString(req.Jsonld)

	err = dbw.InsertData(r.Context(), iri, json.RawMessage(jsonld))
	if err != nil {

		// a duplicate IRI means the exact same data is already stored, so we return the data
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"data_pkey\"") {
			respondJSON(rw, http.StatusOK, NewPostDataResponse(iri, jsonld))
			return
		}

		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostDataResponse(iri, jsonld))
}
