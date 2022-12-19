package db

import (
	"context"
)

// Reader is the interface that wraps database reads.
type Reader interface {

	// GetContent reads content from the database.
	GetContent(ctx context.Context, id int64) (string, error)
}

var _ Reader = &reader{}

type reader struct {
	q *Queries
}

// GetContent reads content from the database.
func (r *reader) GetContent(ctx context.Context, id int64) (string, error) {
	return r.q.GetContent(ctx, id)
}
