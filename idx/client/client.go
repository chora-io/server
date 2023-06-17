package client

import (
	"context"
	"os"

	"github.com/rs/zerolog"

	db "github.com/choraio/server/db/client"
	"github.com/choraio/server/idx/config"
)

// Client is the client.
type Client struct {
	db  db.Database
	log zerolog.Logger
}

// NewClient creates a new client.
func NewClient(cfg config.Config) (Client, error) {
	c := Client{}
	c.log = zerolog.New(os.Stdout)

	// initialize and set client db
	newDb, err := db.NewDatabase(cfg.DatabaseUrl, c.log)
	if err != nil {
		return Client{}, err
	}
	c.db = newDb

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

// GetLastBlock gets the last processed block for a given process.
func (c Client) GetLastBlock(ctx context.Context, chainId, processName string) (int64, error) {
	lastBlock, err := c.db.Reader().GetIdxProcessLastBlock(ctx, chainId, processName)
	if err != nil {
		return 0, err
	}
	return lastBlock, nil
}

// IncrementLastBlock increments the last processed block for a given process.
func (c Client) IncrementLastBlock(ctx context.Context, chainId, processName string) error {
	lastBlock, err := c.GetLastBlock(ctx, chainId, processName)
	if err != nil {
		return err
	}
	err = c.db.Writer().UpdateIdxProcessLastBlock(ctx, chainId, processName, lastBlock+1)
	if err != nil {
		return err
	}
	return nil
}
