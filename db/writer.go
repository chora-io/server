package db

import (
	"context"
)

// Writer is the interface that wraps database writes.
type Writer interface {

	// PostData writes data to the database.
	PostData(ctx context.Context, iri, context, jsonld string) error
}

var _ Writer = &writer{}

type writer struct {
	q *Queries
}

// PostData writes data to the database.
func (w *writer) PostData(ctx context.Context, iri, context, jsonld string) error {
	return w.q.PostData(ctx, PostDataParams{
		Iri:     iri,
		Context: context,
		Jsonld:  jsonld,
	})
}
