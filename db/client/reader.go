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

	// SelectIdxGroupProposal reads data from the database.
	SelectIdxGroupProposal(ctx context.Context, chainId string, proposalId int64) (json.RawMessage, error)

	// SelectIdxGroupProposals reads data from the database.
	SelectIdxGroupProposals(ctx context.Context, chainId string) ([]json.RawMessage, error)

	// SelectIdxProcessLastBlock reads data from the database.
	SelectIdxProcessLastBlock(ctx context.Context, chainId string, processName string) (int64, error)
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

// SelectIdxGroupProposal reads data from the database.
func (r *reader) SelectIdxGroupProposal(ctx context.Context, chainId string, proposalId int64) (json.RawMessage, error) {
	return r.q.SelectIdxGroupProposal(ctx, SelectIdxGroupProposalParams{
		ChainID:    chainId,
		ProposalID: proposalId,
	})
}

// SelectIdxGroupProposals reads data from the database.
func (r *reader) SelectIdxGroupProposals(ctx context.Context, chainId string) ([]json.RawMessage, error) {
	return r.q.SelectIdxGroupProposals(ctx, chainId)
}

// SelectIdxProcessLastBlock reads data from the database.
func (r *reader) SelectIdxProcessLastBlock(ctx context.Context, chainId string, processName string) (int64, error) {
	return r.q.SelectIdxProcessLastBlock(ctx, SelectIdxProcessLastBlockParams{
		ChainID:     chainId,
		ProcessName: processName,
	})
}
