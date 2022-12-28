package db

import (
	"context"
)

// Reader is the interface that wraps database reads.
type Reader interface {

	// GetData reads data from the database.
	GetData(ctx context.Context, iri string) (Datum, error)
}

var _ Reader = &reader{}

type reader struct {
	q *Queries
}

// GetData reads data from the database.
func (r *reader) GetData(ctx context.Context, iri string) (Datum, error) {
	return r.q.GetData(ctx, iri)
}
