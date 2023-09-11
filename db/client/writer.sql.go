// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: writer.sql

package client

import (
	"context"
	"encoding/json"
)

const insertAuthUser = `-- name: InsertAuthUser :exec
insert into auth_user (address, created_at, last_authenticated) values ($1, now(), now())
`

func (q *Queries) InsertAuthUser(ctx context.Context, address string) error {
	_, err := q.db.ExecContext(ctx, insertAuthUser, address)
	return err
}

const insertData = `-- name: InsertData :exec
insert into data (iri, jsonld) values ($1, $2)
`

type InsertDataParams struct {
	Iri    string
	Jsonld json.RawMessage
}

func (q *Queries) InsertData(ctx context.Context, arg InsertDataParams) error {
	_, err := q.db.ExecContext(ctx, insertData, arg.Iri, arg.Jsonld)
	return err
}

const insertIdxGroupProposal = `-- name: InsertIdxGroupProposal :exec
insert into idx_group_proposal (chain_id, proposal_id, group_id, proposal) values ($1, $2, $3, $4)
`

type InsertIdxGroupProposalParams struct {
	ChainID    string
	ProposalID int64
	GroupID    int64
	Proposal   json.RawMessage
}

func (q *Queries) InsertIdxGroupProposal(ctx context.Context, arg InsertIdxGroupProposalParams) error {
	_, err := q.db.ExecContext(ctx, insertIdxGroupProposal,
		arg.ChainID,
		arg.ProposalID,
		arg.GroupID,
		arg.Proposal,
	)
	return err
}

const insertIdxGroupVote = `-- name: InsertIdxGroupVote :exec
insert into idx_group_vote (chain_id, proposal_id, voter, vote) values ($1, $2, $3, $4)
`

type InsertIdxGroupVoteParams struct {
	ChainID    string
	ProposalID int64
	Voter      string
	Vote       json.RawMessage
}

func (q *Queries) InsertIdxGroupVote(ctx context.Context, arg InsertIdxGroupVoteParams) error {
	_, err := q.db.ExecContext(ctx, insertIdxGroupVote,
		arg.ChainID,
		arg.ProposalID,
		arg.Voter,
		arg.Vote,
	)
	return err
}

const insertIdxProcessLastBlock = `-- name: InsertIdxProcessLastBlock :exec
insert into idx_process (chain_id, process_name, last_block) values ($1, $2, $3)
`

type InsertIdxProcessLastBlockParams struct {
	ChainID     string
	ProcessName string
	LastBlock   int64
}

func (q *Queries) InsertIdxProcessLastBlock(ctx context.Context, arg InsertIdxProcessLastBlockParams) error {
	_, err := q.db.ExecContext(ctx, insertIdxProcessLastBlock, arg.ChainID, arg.ProcessName, arg.LastBlock)
	return err
}

const insertIdxSkippedBlock = `-- name: InsertIdxSkippedBlock :exec
insert into idx_skipped_block (chain_id, process_name, skipped_block, reason) values ($1, $2, $3, $4)
`

type InsertIdxSkippedBlockParams struct {
	ChainID      string
	ProcessName  string
	SkippedBlock int64
	Reason       string
}

func (q *Queries) InsertIdxSkippedBlock(ctx context.Context, arg InsertIdxSkippedBlockParams) error {
	_, err := q.db.ExecContext(ctx, insertIdxSkippedBlock,
		arg.ChainID,
		arg.ProcessName,
		arg.SkippedBlock,
		arg.Reason,
	)
	return err
}

const updateAuthUserLastAuthenticated = `-- name: UpdateAuthUserLastAuthenticated :exec
update auth_user set last_authenticated=now() where address=$1
`

func (q *Queries) UpdateAuthUserLastAuthenticated(ctx context.Context, address string) error {
	_, err := q.db.ExecContext(ctx, updateAuthUserLastAuthenticated, address)
	return err
}

const updateIdxGroupProposal = `-- name: UpdateIdxGroupProposal :exec
update idx_group_proposal set proposal=$3 where chain_id=$1 and proposal_id=$2
`

type UpdateIdxGroupProposalParams struct {
	ChainID    string
	ProposalID int64
	Proposal   json.RawMessage
}

func (q *Queries) UpdateIdxGroupProposal(ctx context.Context, arg UpdateIdxGroupProposalParams) error {
	_, err := q.db.ExecContext(ctx, updateIdxGroupProposal, arg.ChainID, arg.ProposalID, arg.Proposal)
	return err
}

const updateIdxGroupVote = `-- name: UpdateIdxGroupVote :exec
update idx_group_vote set vote=$4 where chain_id=$1 and proposal_id=$2 and voter=$3
`

type UpdateIdxGroupVoteParams struct {
	ChainID    string
	ProposalID int64
	Voter      string
	Vote       json.RawMessage
}

func (q *Queries) UpdateIdxGroupVote(ctx context.Context, arg UpdateIdxGroupVoteParams) error {
	_, err := q.db.ExecContext(ctx, updateIdxGroupVote,
		arg.ChainID,
		arg.ProposalID,
		arg.Voter,
		arg.Vote,
	)
	return err
}

const updateIdxProcessLastBlock = `-- name: UpdateIdxProcessLastBlock :exec
update idx_process set last_block=$3 where chain_id=$1 and process_name=$2
`

type UpdateIdxProcessLastBlockParams struct {
	ChainID     string
	ProcessName string
	LastBlock   int64
}

func (q *Queries) UpdateIdxProcessLastBlock(ctx context.Context, arg UpdateIdxProcessLastBlockParams) error {
	_, err := q.db.ExecContext(ctx, updateIdxProcessLastBlock, arg.ChainID, arg.ProcessName, arg.LastBlock)
	return err
}
