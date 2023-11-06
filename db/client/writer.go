package client

import (
	"context"
	"database/sql"
	"encoding/json"
)

// Writer is the interface that wraps database writes.
type Writer interface {

	// auth queries

	// InsertAuthUserWithAddress writes data to the database.
	InsertAuthUserWithAddress(ctx context.Context, address string) error

	// InsertAuthUserWithEmail writes data to the database.
	InsertAuthUserWithEmail(ctx context.Context, email string) error

	// InsertAuthUserWithUsername writes data to the database.
	InsertAuthUserWithUsername(ctx context.Context, username string) error

	// UpdateAuthUserAddress writes data to the database.
	UpdateAuthUserAddress(ctx context.Context, userId string, address string) error

	// UpdateAuthUserEmail writes data to the database.
	UpdateAuthUserEmail(ctx context.Context, userId string, email string) error

	// UpdateAuthUserUsername writes data to the database.
	UpdateAuthUserUsername(ctx context.Context, userId string, username string) error

	// data queries

	// InsertData writes data to the database.
	InsertData(ctx context.Context, iri string, jsonld json.RawMessage) error

	// indexer queries

	// InsertIdxGroupProposal writes data to the database.
	InsertIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, groupId int64, proposal json.RawMessage) error

	// InsertIdxGroupVote writes data to the database.
	InsertIdxGroupVote(ctx context.Context, chainId string, proposalId int64, voter string, vote json.RawMessage) error

	// InsertIdxProcessLastBlock writes data to the database.
	InsertIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlack int64) error

	// InsertIdxSkippedBlock writes data to the database.
	InsertIdxSkippedBlock(ctx context.Context, chainId string, processName string, skippedBlock int64, reason string) error

	// UpdateIdxGroupProposal writes data to the database.
	UpdateIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error

	// UpdateIdxGroupVote writes data to the database.
	UpdateIdxGroupVote(ctx context.Context, chainId string, proposalId int64, voter string, vote json.RawMessage) error

	// UpdateIdxProcessLastBlock writes data to the database.
	UpdateIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlock int64) error
}

var _ Writer = &writer{}

type writer struct {
	q *Queries
}

// auth queries

func (w *writer) InsertAuthUserWithAddress(ctx context.Context, address string) error {
	return w.q.InsertAuthUserWithAddress(ctx, sql.NullString{
		String: address,
		Valid:  len(address) > 0,
	})
}

func (w *writer) InsertAuthUserWithEmail(ctx context.Context, email string) error {
	return w.q.InsertAuthUserWithEmail(ctx, sql.NullString{
		String: email,
		Valid:  len(email) > 0,
	})
}

func (w *writer) InsertAuthUserWithUsername(ctx context.Context, username string) error {
	return w.q.InsertAuthUserWithUsername(ctx, sql.NullString{
		String: username,
		Valid:  len(username) > 0,
	})
}

func (w *writer) UpdateAuthUserAddress(ctx context.Context, userId string, address string) error {
	return w.q.UpdateAuthUserAddress(ctx, UpdateAuthUserAddressParams{
		ID: userId,
		Address: sql.NullString{
			String: address,
			Valid:  len(address) > 0,
		},
	})
}

func (w *writer) UpdateAuthUserEmail(ctx context.Context, userId string, email string) error {
	return w.q.UpdateAuthUserEmail(ctx, UpdateAuthUserEmailParams{
		ID: userId,
		Email: sql.NullString{
			String: email,
			Valid:  len(email) > 0,
		},
	})
}

func (w *writer) UpdateAuthUserUsername(ctx context.Context, userId string, username string) error {
	return w.q.UpdateAuthUserUsername(ctx, UpdateAuthUserUsernameParams{
		ID: userId,
		Username: sql.NullString{
			String: username,
			Valid:  len(username) > 0,
		},
	})
}

// data queries

func (w *writer) InsertData(ctx context.Context, iri string, jsonld json.RawMessage) error {
	return w.q.InsertData(ctx, InsertDataParams{
		Iri:    iri,
		Jsonld: jsonld,
	})
}

// indexer queries

func (w *writer) InsertIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, groupId int64, proposal json.RawMessage) error {
	return w.q.InsertIdxGroupProposal(ctx, InsertIdxGroupProposalParams{
		ChainID:    chainId,
		ProposalID: proposalId,
		GroupID:    groupId,
		Proposal:   proposal,
	})
}

func (w *writer) InsertIdxGroupVote(ctx context.Context, chainId string, proposalId int64, voter string, vote json.RawMessage) error {
	return w.q.InsertIdxGroupVote(ctx, InsertIdxGroupVoteParams{
		ChainID:    chainId,
		ProposalID: proposalId,
		Voter:      voter,
		Vote:       vote,
	})
}

func (w *writer) InsertIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlack int64) error {
	return w.q.InsertIdxProcessLastBlock(ctx, InsertIdxProcessLastBlockParams{
		ChainID:     chainId,
		ProcessName: processName,
		LastBlock:   lastBlack,
	})
}

func (w *writer) InsertIdxSkippedBlock(ctx context.Context, chainId string, processName string, skippedBlock int64, reason string) error {
	return w.q.InsertIdxSkippedBlock(ctx, InsertIdxSkippedBlockParams{
		ChainID:      chainId,
		ProcessName:  processName,
		SkippedBlock: skippedBlock,
		Reason:       reason,
	})
}

func (w *writer) UpdateIdxGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error {
	return w.q.UpdateIdxGroupProposal(ctx, UpdateIdxGroupProposalParams{
		ChainID:    chainId,
		ProposalID: proposalId,
		Proposal:   proposal,
	})
}

func (w *writer) UpdateIdxGroupVote(ctx context.Context, chainId string, proposalId int64, voter string, vote json.RawMessage) error {
	return w.q.UpdateIdxGroupVote(ctx, UpdateIdxGroupVoteParams{
		ChainID:    chainId,
		ProposalID: proposalId,
		Voter:      voter,
		Vote:       vote,
	})
}

func (w *writer) UpdateIdxProcessLastBlock(ctx context.Context, chainId string, processName string, lastBlock int64) error {
	return w.q.UpdateIdxProcessLastBlock(ctx, UpdateIdxProcessLastBlockParams{
		ChainID:     chainId,
		ProcessName: processName,
		LastBlock:   lastBlock,
	})
}
