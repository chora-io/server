package app

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/choraio/server/db/client"
	"github.com/gorilla/mux"
)

func GetIdxGroupProposal(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	chainId := vars["chain_id"]
	proposalId := vars["proposal_id"]

	parsedId, err := strconv.ParseInt(proposalId, 0, 64)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
	}

	d, err := dbr.SelectIdxGroupProposal(r.Context(), chainId, parsedId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(
				rw,
				http.StatusNotFound,
				fmt.Sprintf("proposal with id %s does not exist", proposalId),
			)
		} else {
			respondError(rw, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondJSON(rw, http.StatusOK, NewGetIdxGroupProposalResponse(d))
}

func GetIdxGroupProposals(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	chainId := vars["chain_id"]

	d, err := dbr.SelectIdxGroupProposals(r.Context(), chainId)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
	}

	respondJSON(rw, http.StatusOK, NewGetIdxGroupProposalsResponse(d))
}
