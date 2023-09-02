package app

import "encoding/json"

type GetIdxGroupProposalResponse struct {
	// Proposal is the proposal.
	Proposal json.RawMessage `json:"proposal"`
}

func NewGetIdxGroupProposalResponse(proposal json.RawMessage) GetIdxGroupProposalResponse {
	return GetIdxGroupProposalResponse{
		Proposal: proposal,
	}
}

type GetIdxGroupProposalsResponse struct {
	// Proposals are the proposals.
	Proposals []json.RawMessage `json:"proposals"`
}

func NewGetIdxGroupProposalsResponse(proposals []json.RawMessage) GetIdxGroupProposalsResponse {
	return GetIdxGroupProposalsResponse{
		Proposals: proposals,
	}
}

type GetIdxGroupVoteResponse struct {
	// Vote is the vote.
	Vote json.RawMessage `json:"vote"`
}

func NewGetIdxGroupVoteResponse(vote json.RawMessage) GetIdxGroupVoteResponse {
	return GetIdxGroupVoteResponse{
		Vote: vote,
	}
}

type GetIdxGroupVotesResponse struct {
	// Votes are the votes.
	Votes []json.RawMessage `json:"votes"`
}

func NewGetIdxGroupVotesResponse(votes []json.RawMessage) GetIdxGroupVotesResponse {
	return GetIdxGroupVotesResponse{
		Votes: votes,
	}
}
