package process

import (
	"context"
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

	// StartBlock is the overriding start block height from which the process will start. This forces each process
	// to start at a specific block height. If not set, each process will either start from the last processed block
	// or from the first block (i.e. block 1) if no last processed block exists.
	StartBlock int64
}

func (p Params) ValidateParams() error {
	if p.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if p.ChainId == "" {
		return fmt.Errorf("chain id cannot be empty")
	}

	if p.StartBlock < 0 {
		return fmt.Errorf("start block cannot be negative")
	}

	return nil
}
