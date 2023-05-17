// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: writer.sql

package client

import (
	"context"
	"encoding/json"
)

const postData = `-- name: PostData :exec
insert into data (iri, jsonld) values ($1, $2)
`

type PostDataParams struct {
	Iri    string
	Jsonld json.RawMessage
}

func (q *Queries) PostData(ctx context.Context, arg PostDataParams) error {
	_, err := q.db.ExecContext(ctx, postData, arg.Iri, arg.Jsonld)
	return err
}