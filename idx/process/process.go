package process

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/choraio/server/idx/client"
)

// Function is the function used to advance the process.
type Function func(ctx context.Context, p Params) error

// Params are the process parameters.
type Params struct {
	// Name is the name of the process.
	Name string

	// ChainId is the chain id of the network (e.g. chora-testnet-1, regen-redwood-1).
	ChainId string

	// Client is the client that wraps the database and connects to the network.
	Client client.Client

	// StartBlock is the starting block height used when starting a new process (default 1).
	StartBlock int64
}

func (p Params) ValidateParams() error {
	if p.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if p.ChainId == "" {
		return fmt.Errorf("chain id cannot be empty")
	}

	if p.StartBlock < 1 {
		return fmt.Errorf("start block must be a positive integer")
	}

	return nil
}

// AdvanceProcess determines the last block and next block for the process
func AdvanceProcess(ctx context.Context, p Params) (int64, int64, error) {

	// get latest block from connected network (chain)
	latestBlock, err := p.Client.GetLatestBlockHeight()
	if err != nil {
		return 0, 0, err
	}

	// select last processed block for chain and process from database
	lastBlock, err := p.Client.SelectProcessLastBlock(ctx, p.ChainId, p.Name)

	// handle new process and process error
	if err == sql.ErrNoRows {
		fmt.Println(p.Name, "last processed block does not exist")

		// set last processed block to start block - 1
		lastBlock = p.StartBlock - 1

		fmt.Println(p.Name, "inserting last processed block", lastBlock)

		// insert last processed block
		err = p.Client.InsertProcessLastBlock(ctx, p.ChainId, p.Name, lastBlock)
		if err != nil {
			return 0, 0, err
		}
	} else if err != nil {
		return 0, 0, err
	}

	// handle process in sync
	if lastBlock == latestBlock {
		fmt.Println(p.Name, "process in sync", lastBlock, latestBlock)

		// return matching last block and latest block (next block)
		return lastBlock, latestBlock, nil
	}

	// handle process mismatch
	if latestBlock < lastBlock {
		fmt.Println(p.Name, "updating last processed block to latest block")

		// set last processed block to latest block
		lastBlock = latestBlock

		// update last processed block to latest block
		err = p.Client.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, latestBlock)
		if err != nil {
			return 0, 0, err
		}
	}

	fmt.Println(p.Name, "last block", lastBlock)

	nextBlock := lastBlock + 1

	fmt.Println(p.Name, "next block", nextBlock)

	return lastBlock, nextBlock, nil
}
