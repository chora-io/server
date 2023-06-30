package client

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/x/group"
	"github.com/rs/zerolog"

	db "github.com/choraio/server/db/client"
	"github.com/choraio/server/idx/client/cosmos"
)

// Client is the client.
type Client struct {
	db  db.Database
	cc  cosmos.Client
	log zerolog.Logger
}

// NewClient wraps the provided database and logger and creates a new cosmos client.
func NewClient(rpcUrl string, db db.Database, log zerolog.Logger) (Client, error) {
	// wrap db and logger
	c := Client{
		db:  db,
		log: log,
	}

	// initialize and set cosmos client
	newCosmos, err := cosmos.NewClient(rpcUrl)
	if err != nil {
		return Client{}, err
	}
	c.cc = newCosmos

	return c, nil
}

// Close shuts down the client.
func (c Client) Close() error {

	// close cosmos client
	err := c.cc.Close()
	if err != nil {
		return err
	}

	return nil
}

// database queries

// InsertGroupProposal adds a group proposal to the database.
func (c Client) InsertGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error {
	err := c.db.Writer().InsertIdxGroupProposal(ctx, chainId, proposalId, proposal)
	if err != nil {
		return err
	}
	return nil
}

// InsertProcessLastBlock adds a process with block height to the database.
func (c Client) InsertProcessLastBlock(ctx context.Context, chainId, processName string, lastBlock int64) error {
	err := c.db.Writer().InsertIdxProcessLastBlock(ctx, chainId, processName, lastBlock)
	if err != nil {
		return err
	}
	return nil
}

// SelectProcessLastBlock gets the last processed block for a given process.
func (c Client) SelectProcessLastBlock(ctx context.Context, chainId, processName string) (int64, error) {
	lastBlock, err := c.db.Reader().SelectIdxProcessLastBlock(ctx, chainId, processName)
	if err != nil {
		return 0, err
	}
	return lastBlock, nil
}

// UpdateGroupProposal updates a group proposal in the database.
func (c Client) UpdateGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error {
	err := c.db.Writer().UpdateIdxGroupProposal(ctx, chainId, proposalId, proposal)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProcessLastBlock updates the last processed block for a given process.
func (c Client) UpdateProcessLastBlock(ctx context.Context, chainId, processName string, lastBlock int64) error {
	err := c.db.Writer().UpdateIdxProcessLastBlock(ctx, chainId, processName, lastBlock)
	if err != nil {
		return err
	}
	return nil
}

// cosmos blockchain queries

// GetGroupEventProposalPruned gets any array of group.v1.EventProposalPruned from block height.
func (c Client) GetGroupEventProposalPruned(height int64) ([]group.EventProposalPruned, error) {
	return c.cc.GetGroupEventProposalPruned(height)
}

// GetGroupProposal gets a group proposal by proposal id at block height.
func (c Client) GetGroupProposal(height int64, proposalId int64) (json.RawMessage, error) {
	return c.cc.GetGroupProposal(height, proposalId)
}

// GetLatestBlockHeight gets the latest block height.
func (c Client) GetLatestBlockHeight() (int64, error) {
	return c.cc.GetLatestBlockHeight()
}
