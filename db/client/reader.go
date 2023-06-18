package client

import (
	"context"
	"encoding/json"
)

// Reader is the interface that wraps database reads.
type Reader interface {

	// data queries

	// GetData reads data from the database.
	GetData(ctx context.Context, iri string) (Datum, error)

	// indexer queries

	// GetIdxGroupProposal reads data from the database.
	GetIdxGroupProposal(ctx context.Context, chainId string, proposalId int64) (json.RawMessage, error)

	// GetIdxGroupProposals reads data from the database.
	GetIdxGroupProposals(ctx context.Context, chainId string) ([]json.RawMessage, error)

	// GetIdxProcessLastBlock reads data from the database.
	GetIdxProcessLastBlock(ctx context.Context, chainId string, processName string) (int64, error)
}

var _ Reader = &reader{}

type reader struct {
	q *Queries
}

// data queries

// GetData reads data from the database.
func (r *reader) GetData(ctx context.Context, iri string) (Datum, error) {
	return r.q.GetData(ctx, iri)
}

// indexer queries

// GetIdxGroupProposal reads data from the database.
func (r *reader) GetIdxGroupProposal(ctx context.Context, chainId string, proposalId int64) (json.RawMessage, error) {
	return r.q.GetIdxGroupProposal(ctx, GetIdxGroupProposalParams{
		ChainID:    chainId,
		ProposalID: proposalId,
	})
}

// GetIdxGroupProposals reads data from the database.
func (r *reader) GetIdxGroupProposals(ctx context.Context, chainId string) ([]json.RawMessage, error) {
	return r.q.GetIdxGroupProposals(ctx, chainId)
}

// GetIdxProcessLastBlock reads data from the database.
func (r *reader) GetIdxProcessLastBlock(ctx context.Context, chainId string, processName string) (int64, error) {
	return r.q.GetIdxProcessLastBlock(ctx, GetIdxProcessLastBlockParams{
		ChainID:     chainId,
		ProcessName: processName,
	})
}
