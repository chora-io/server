package client

import (
	"context"
	"encoding/json"
)

// Writer is the interface that wraps database writes.
type Writer interface {

	// PostData writes data to the database.
	PostData(ctx context.Context, iri string, jsonld json.RawMessage) error

	// AddIdxGroupProposal writes data to the database.
	AddIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error

	// UpdateIdxGroupProposal writes data to the database.
	UpdateIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error

	// UpdateIdxProcessLastBlock writes data to the database.
	UpdateIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlock int64) error
}

var _ Writer = &writer{}

type writer struct {
	q *Queries
}

func (w *writer) PostData(ctx context.Context, iri string, jsonld json.RawMessage) error {
	return w.q.PostData(ctx, PostDataParams{
		Iri:    iri,
		Jsonld: jsonld,
	})
}

func (w *writer) AddIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error {
	return w.q.AddIdxGroupProposal(ctx, AddIdxGroupProposalParams{
		ChainID:    chainId,
		ProposalID: proposalId,
		Proposal:   proposal,
	})
}

func (w *writer) UpdateIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error {
	return w.q.UpdateIdxGroupProposal(ctx, UpdateIdxGroupProposalParams{
		ChainID:    chainId,
		ProposalID: proposalId,
		Proposal:   proposal,
	})
}

func (w *writer) UpdateIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlock int64) error {
	return w.q.UpdateIdxProcessLastBlock(ctx, UpdateIdxProcessLastBlockParams{
		ChainID:     chainId,
		ProcessName: processName,
		LastBlock:   lastBlock,
	})
}
