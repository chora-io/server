// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package client

import (
	"encoding/json"
)

// the data table stores linked data
type Datum struct {
	Iri    string
	Jsonld json.RawMessage
}