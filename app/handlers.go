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

	// TODO: propose module without sdk dependencies
	"github.com/regen-network/regen-ledger/x/data/v2"

	"github.com/choraio/server/db"
)

type GetRequestHandlerFunction func(dbr db.Reader, rw http.ResponseWriter, r *http.Request)

func GetData(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	iri := vars["iri"]

	_, err := data.ParseIRI(iri)
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

	respondJSON(rw, http.StatusOK, NewGetDataResponse(iri, d.Context, d.Jsonld))
}

type PostRequestHandlerFunction func(dbw db.Writer, rw http.ResponseWriter, r *http.Request)

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

	ch := data.ContentHash_Graph{
		Hash:                      hash.Sum(nil),
		DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
		CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		MerkleTree:                data.GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED,
	}

	iri, err := ch.ToIRI(&data.IRIOptions{Prefix: "chora"})
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	context := strings.TrimSpace(req.Context)
	jsonld := compactJSONString(req.Jsonld)

	err = dbw.PostData(r.Context(), iri, context, jsonld)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostDataResponse(iri, context, jsonld))
}
