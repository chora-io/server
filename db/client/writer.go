package client

import (
	"context"
	"encoding/json"
)

// Writer is the interface that wraps database writes.
type Writer interface {

	// PostData writes data to the database.
	PostData(ctx context.Context, iri string, jsonld json.RawMessage) error
}

var _ Writer = &writer{}

type writer struct {
	q *Queries
}

// PostData writes data to the database.
func (w *writer) PostData(ctx context.Context, iri string, jsonld json.RawMessage) error {
	return w.q.PostData(ctx, PostDataParams{
		Iri:    iri,
		Jsonld: jsonld,
	})
}
