package app

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	db "github.com/chora-io/server/db/client"
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
	groupId := vars["group_id"]

	parsedId, err := strconv.ParseInt(groupId, 0, 64)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
	}

	d, err := dbr.SelectIdxGroupProposals(r.Context(), chainId, parsedId)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
	}

	respondJSON(rw, http.StatusOK, NewGetIdxGroupProposalsResponse(d))
}

func GetIdxGroupVote(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	chainId := vars["chain_id"]
	proposalId := vars["proposal_id"]
	voter := vars["voter"]

	parsedId, err := strconv.ParseInt(proposalId, 0, 64)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
	}

	d, err := dbr.SelectIdxGroupVote(r.Context(), chainId, parsedId, voter)
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

	respondJSON(rw, http.StatusOK, NewGetIdxGroupVoteResponse(d))
}

func GetIdxGroupVotes(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	chainId := vars["chain_id"]
	proposalId := vars["proposal_id"]

	parsedId, err := strconv.ParseInt(proposalId, 0, 64)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
	}

	d, err := dbr.SelectIdxGroupVotes(r.Context(), chainId, parsedId)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
	}

	respondJSON(rw, http.StatusOK, NewGetIdxGroupVotesResponse(d))
}
