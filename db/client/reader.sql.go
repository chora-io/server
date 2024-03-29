// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: reader.sql

package client

import (
	"context"
	"database/sql"
	"encoding/json"
)

const selectAuthUser = `-- name: SelectAuthUser :one
select id, address, email, username, created_at from auth_user where id=$1
`

func (q *Queries) SelectAuthUser(ctx context.Context, id string) (AuthUser, error) {
	row := q.db.QueryRowContext(ctx, selectAuthUser, id)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Address,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const selectAuthUserByAddress = `-- name: SelectAuthUserByAddress :one
select id, address, email, username, created_at from auth_user where address=$1
`

func (q *Queries) SelectAuthUserByAddress(ctx context.Context, address sql.NullString) (AuthUser, error) {
	row := q.db.QueryRowContext(ctx, selectAuthUserByAddress, address)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Address,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const selectAuthUserByEmail = `-- name: SelectAuthUserByEmail :one
select id, address, email, username, created_at from auth_user where email=$1
`

func (q *Queries) SelectAuthUserByEmail(ctx context.Context, email sql.NullString) (AuthUser, error) {
	row := q.db.QueryRowContext(ctx, selectAuthUserByEmail, email)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Address,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const selectAuthUserByUsername = `-- name: SelectAuthUserByUsername :one
select id, address, email, username, created_at from auth_user where username=$1
`

func (q *Queries) SelectAuthUserByUsername(ctx context.Context, username sql.NullString) (AuthUser, error) {
	row := q.db.QueryRowContext(ctx, selectAuthUserByUsername, username)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.Address,
		&i.Email,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const selectData = `-- name: SelectData :one
select iri, jsonld from data where iri=$1
`

func (q *Queries) SelectData(ctx context.Context, iri string) (Datum, error) {
	row := q.db.QueryRowContext(ctx, selectData, iri)
	var i Datum
	err := row.Scan(&i.Iri, &i.Jsonld)
	return i, err
}

const selectIdxGroupProposal = `-- name: SelectIdxGroupProposal :one
select proposal from idx_group_proposal where chain_id=$1 and proposal_id=$2
`

type SelectIdxGroupProposalParams struct {
	ChainID    string
	ProposalID int64
}

func (q *Queries) SelectIdxGroupProposal(ctx context.Context, arg SelectIdxGroupProposalParams) (json.RawMessage, error) {
	row := q.db.QueryRowContext(ctx, selectIdxGroupProposal, arg.ChainID, arg.ProposalID)
	var proposal json.RawMessage
	err := row.Scan(&proposal)
	return proposal, err
}

const selectIdxGroupProposals = `-- name: SelectIdxGroupProposals :many
select proposal from idx_group_proposal where chain_id=$1 and group_id=$2
`

type SelectIdxGroupProposalsParams struct {
	ChainID string
	GroupID int64
}

func (q *Queries) SelectIdxGroupProposals(ctx context.Context, arg SelectIdxGroupProposalsParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, selectIdxGroupProposals, arg.ChainID, arg.GroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var proposal json.RawMessage
		if err := rows.Scan(&proposal); err != nil {
			return nil, err
		}
		items = append(items, proposal)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectIdxGroupVote = `-- name: SelectIdxGroupVote :one
select vote from idx_group_vote where chain_id=$1 and proposal_id=$2 and voter=$3
`

type SelectIdxGroupVoteParams struct {
	ChainID    string
	ProposalID int64
	Voter      string
}

func (q *Queries) SelectIdxGroupVote(ctx context.Context, arg SelectIdxGroupVoteParams) (json.RawMessage, error) {
	row := q.db.QueryRowContext(ctx, selectIdxGroupVote, arg.ChainID, arg.ProposalID, arg.Voter)
	var vote json.RawMessage
	err := row.Scan(&vote)
	return vote, err
}

const selectIdxGroupVotes = `-- name: SelectIdxGroupVotes :many
select vote from idx_group_vote where chain_id=$1 and proposal_id=$2
`

type SelectIdxGroupVotesParams struct {
	ChainID    string
	ProposalID int64
}

func (q *Queries) SelectIdxGroupVotes(ctx context.Context, arg SelectIdxGroupVotesParams) ([]json.RawMessage, error) {
	rows, err := q.db.QueryContext(ctx, selectIdxGroupVotes, arg.ChainID, arg.ProposalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []json.RawMessage
	for rows.Next() {
		var vote json.RawMessage
		if err := rows.Scan(&vote); err != nil {
			return nil, err
		}
		items = append(items, vote)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectIdxProcessLastBlock = `-- name: SelectIdxProcessLastBlock :one
select last_block from idx_process where chain_id=$1 and process_name=$2
`

type SelectIdxProcessLastBlockParams struct {
	ChainID     string
	ProcessName string
}

func (q *Queries) SelectIdxProcessLastBlock(ctx context.Context, arg SelectIdxProcessLastBlockParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, selectIdxProcessLastBlock, arg.ChainID, arg.ProcessName)
	var last_block int64
	err := row.Scan(&last_block)
	return last_block, err
}
