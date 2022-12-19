package db

import (
	"context"
)

// Writer is the interface that wraps database writes.
type Writer interface {

	// PostContent writes content to the database.
	PostContent(ctx context.Context, content string) (int64, error)
}

var _ Writer = &writer{}

type writer struct {
	q *Queries
}

// PostContent writes content to the database.
func (w *writer) PostContent(ctx context.Context, body string) (int64, error) {
	return w.q.PostContent(ctx, body)
}
