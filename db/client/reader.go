package client

import (
	"context"
	"database/sql"
	"encoding/json"
)

// Reader is the interface that wraps database reads.
type Reader interface {

	// auth queries

	// SelectAuthUser reads data from the database.
	SelectAuthUser(ctx context.Context, userId string) (AuthUser, error)

	// SelectAuthUserByAddress reads data from the database.
	SelectAuthUserByAddress(ctx context.Context, address string) (AuthUser, error)

	// SelectAuthUserByEmail reads data from the database.
	SelectAuthUserByEmail(ctx context.Context, email string) (AuthUser, error)

	// SelectAuthUserByUsername reads data from the database.
	SelectAuthUserByUsername(ctx context.Context, username string) (AuthUser, error)

	// data queries

	// GetData reads data from the database.
	GetData(ctx context.Context, iri string) (Datum, error)

	// indexer queries

	// SelectIdxGroupProposal reads data from the database.
	SelectIdxGroupProposal(ctx context.Context, chainId string, proposalId int64) (json.RawMessage, error)

	// SelectIdxGroupProposals reads data from the database.
	SelectIdxGroupProposals(ctx context.Context, chainId string, groupId int64) ([]json.RawMessage, error)

	// SelectIdxGroupVote reads data from the database.
	SelectIdxGroupVote(ctx context.Context, chainId string, proposalId int64, vote string) (json.RawMessage, error)

	// SelectIdxGroupVotes reads data from the database.
	SelectIdxGroupVotes(ctx context.Context, chainId string, proposalId int64) ([]json.RawMessage, error)

	// SelectIdxProcessLastBlock reads data from the database.
	SelectIdxProcessLastBlock(ctx context.Context, chainId string, processName string) (int64, error)
}

var _ Reader = &reader{}

type reader struct {
	q *Queries
}

// GroupProposalParams is used to select proposals by group_id.
type GroupProposalParams struct {
	GroupID string `json:"group_id"`
}

// auth queries

// SelectAuthUser reads data from the database.
func (r *reader) SelectAuthUser(ctx context.Context, userId string) (AuthUser, error) {
	return r.q.SelectAuthUser(ctx, userId)
}

// SelectAuthUserByAddress reads data from the database.
func (r *reader) SelectAuthUserByAddress(ctx context.Context, address string) (AuthUser, error) {
	return r.q.SelectAuthUserByAddress(ctx, sql.NullString{
		String: address,
		Valid:  len(address) > 0,
	})
}

// SelectAuthUserByEmail reads data from the database.
func (r *reader) SelectAuthUserByEmail(ctx context.Context, email string) (AuthUser, error) {
	return r.q.SelectAuthUserByEmail(ctx, sql.NullString{
		String: email,
		Valid:  len(email) > 0,
	})
}

// SelectAuthUserByUsername reads data from the database.
func (r *reader) SelectAuthUserByUsername(ctx context.Context, username string) (AuthUser, error) {
	return r.q.SelectAuthUserByUsername(ctx, sql.NullString{
		String: username,
		Valid:  len(username) > 0,
	})
}

// data queries

// GetData reads data from the database.
func (r *reader) GetData(ctx context.Context, iri string) (Datum, error) {
	return r.q.SelectData(ctx, iri)
}

// indexer queries

// SelectIdxGroupProposal reads data from the database.
func (r *reader) SelectIdxGroupProposal(ctx context.Context, chainId string, proposalId int64) (json.RawMessage, error) {
	return r.q.SelectIdxGroupProposal(ctx, SelectIdxGroupProposalParams{
		ChainID:    chainId,
		ProposalID: proposalId,
	})
}

// SelectIdxGroupProposals reads data from the database.
func (r *reader) SelectIdxGroupProposals(ctx context.Context, chainId string, groupId int64) ([]json.RawMessage, error) {
	return r.q.SelectIdxGroupProposals(ctx, SelectIdxGroupProposalsParams{
		ChainID: chainId,
		GroupID: groupId,
	})
}

// SelectIdxGroupVote reads data from the database.
func (r *reader) SelectIdxGroupVote(ctx context.Context, chainId string, proposalId int64, voter string) (json.RawMessage, error) {
	return r.q.SelectIdxGroupVote(ctx, SelectIdxGroupVoteParams{
		ChainID:    chainId,
		ProposalID: proposalId,
		Voter:      voter,
	})
}

// SelectIdxGroupVotes reads data from the database.
func (r *reader) SelectIdxGroupVotes(ctx context.Context, chainId string, proposalId int64) ([]json.RawMessage, error) {
	return r.q.SelectIdxGroupVotes(ctx, SelectIdxGroupVotesParams{
		ChainID:    chainId,
		ProposalID: proposalId,
	})
}

// SelectIdxProcessLastBlock reads data from the database.
func (r *reader) SelectIdxProcessLastBlock(ctx context.Context, chainId string, processName string) (int64, error) {
	return r.q.SelectIdxProcessLastBlock(ctx, SelectIdxProcessLastBlockParams{
		ChainID:     chainId,
		ProcessName: processName,
	})
}
