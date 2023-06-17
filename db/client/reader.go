package client

import (
	"context"
)

// Reader is the interface that wraps database reads.
type Reader interface {

	// GetData reads data from the database.
	GetData(ctx context.Context, iri string) (Datum, error)

	// GetIdxProcessLastBlock reads data from the database.
	GetIdxProcessLastBlock(ctx context.Context, chainId string, processName string) (int64, error)
}

var _ Reader = &reader{}

type reader struct {
	q *Queries
}

// GetData reads data from the database.
func (r *reader) GetData(ctx context.Context, iri string) (Datum, error) {
	return r.q.GetData(ctx, iri)
}

// GetIdxProcessLastBlock reads data from the database.
func (r *reader) GetIdxProcessLastBlock(ctx context.Context, chainId string, processName string) (int64, error) {
	return r.q.GetIdxProcessLastBlock(ctx, GetIdxProcessLastBlockParams{
		ChainID:     chainId,
		ProcessName: processName,
	})
}
