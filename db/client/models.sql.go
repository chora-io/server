// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package client

import (
	"encoding/json"
	"time"
)

// authenticated user
type AuthUser struct {
	Address           string
	CreatedAt         time.Time
	LastAuthenticated time.Time
}

// the data table stores linked data
type Datum struct {
	Iri    string
	Jsonld json.RawMessage
}

// the final state of a group proposal for a given chain
type IdxGroupProposal struct {
	ChainID    string
	ProposalID int64
	GroupID    int64
	Proposal   json.RawMessage
}

// the final state of a group vote for a given chain
type IdxGroupVote struct {
	ChainID    string
	ProposalID int64
	Voter      string
	Vote       json.RawMessage
}

// the idx process table stores information about a process
type IdxProcess struct {
	ChainID     string
	ProcessName string
	LastBlock   int64
}
