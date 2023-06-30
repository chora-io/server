package client

import (
	"context"
	"encoding/json"
)

// Writer is the interface that wraps database writes.
type Writer interface {

	// data queries

	// InsertData writes data to the database.
	InsertData(ctx context.Context, iri string, jsonld json.RawMessage) error

	// indexer queries

	// InsertIdxGroupProposal writes data to the database.
	InsertIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error

	// InsertIdxProcessLastBlock writes data to the database.
	InsertIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlack int64) error

	// UpdateIdxGroupProposal writes data to the database.
	UpdateIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error

	// UpdateIdxProcessLastBlock writes data to the database.
	UpdateIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlock int64) error
}

var _ Writer = &writer{}

type writer struct {
	q *Queries
}

// data queries

func (w *writer) InsertData(ctx context.Context, iri string, jsonld json.RawMessage) error {
	return w.q.InsertData(ctx, InsertDataParams{
		Iri:    iri,
		Jsonld: jsonld,
	})
}

// indexer queries

func (w *writer) InsertIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error {
	return w.q.InsertIdxGroupProposal(ctx, InsertIdxGroupProposalParams{
		ChainID:    chainId,
		ProposalID: proposalId,
		Proposal:   proposal,
	})
}

func (w *writer) InsertIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlack int64) error {
	return w.q.InsertIdxProcessLastBlock(ctx, InsertIdxProcessLastBlockParams{
		ChainID:     chainId,
		ProcessName: processName,
		LastBlock:   lastBlack,
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
