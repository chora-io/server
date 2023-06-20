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
