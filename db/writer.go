package db

import (
	"context"
)

// Writer is the interface that wraps database writes.
type Writer interface {

	// PostData writes data to the database.
	PostData(ctx context.Context, canon string, context string, jsonld string) (int32, error)
}

var _ Writer = &writer{}

type writer struct {
	q *Queries
}

// PostData writes data to the database.
func (w *writer) PostData(ctx context.Context, canon string, context string, jsonld string) (int32, error) {
	return w.q.PostData(ctx, PostDataParams{
		Canon:   canon,
		Context: context,
		Jsonld:  jsonld,
	})
}
