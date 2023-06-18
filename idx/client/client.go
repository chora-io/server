package client

import (
	"context"
	"encoding/json"
	"os"

	"github.com/rs/zerolog"

	db "github.com/choraio/server/db/client"
	"github.com/choraio/server/idx/client/cosmos"
	"github.com/choraio/server/idx/config"
)

// Client is the client.
type Client struct {
	db  db.Database
	cc  cosmos.Client
	log zerolog.Logger

	// ChainId is the chain id from the configuration. Information about each process such as last
	// processed block is stored in the database using (chain id, process name). Indexed blockchain
	// state is stored using only the chain id and therefore is not specific to a process.
	ChainId string

	// ProcessName is the name of the process. Information about each process such as last processed
	// block is stored in the database using (chain id, process name) as the primary key.
	ProcessName string
}

// NewClient creates a new client.
func NewClient(cfg config.Config) (Client, error) {
	c := Client{}

	// set logger
	c.log = zerolog.New(os.Stdout)

	// initialize and set db client
	newDb, err := db.NewDatabase(cfg.DatabaseUrl, c.log)
	if err != nil {
		return Client{}, err
	}
	c.db = newDb

	// initialize and set cosmos client
	newCosmos, err := cosmos.NewClient(cfg.ChainRpc)
	if err != nil {
		return Client{}, err
	}
	c.cc = newCosmos

	return c, nil
}

// Close closes the client.
func (c Client) Close() error {
	if c.db != nil {
		err := c.db.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// database queries

// AddGroupProposal adds a group proposal to the database.
func (c Client) AddGroupProposal(ctx context.Context, chainId string, proposalId int64, proposal json.RawMessage) error {
	err := c.db.Writer().AddIdxGroupProposal(ctx, chainId, proposalId, proposal)
	if err != nil {
		return err
	}
	return nil
}

// GetProcessLastBlock gets the last processed block for a given process.
func (c Client) GetProcessLastBlock(ctx context.Context, chainId, processName string) (int64, error) {
	lastBlock, err := c.db.Reader().GetIdxProcessLastBlock(ctx, chainId, processName)
	if err != nil {
		return 0, err
	}
	return lastBlock, nil
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

// GetGroupEventProposalPruned gets group.v1.EventProposalPruned from the block.
func (c Client) GetGroupEventProposalPruned(height int64) ([]json.RawMessage, error) {
	return c.cc.GetGroupEventProposalPruned(height)
}

// GetGroupProposal gets a group proposal by proposal id at a given block height.
func (c Client) GetGroupProposal(height int64, proposalId int64) (json.RawMessage, error) {
	return c.cc.GetGroupProposal(height, proposalId)
}
