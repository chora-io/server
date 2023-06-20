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

	// ChainRpc is the rpc endpoint for the network (e.g. testnet.chora.io:9090, redwood.chora.io:9090).
	//
	// TODO: When ChainRpc is empty, the Client will be used. When ChainRpc is provided, a new client
	// will be created from the rpc url and Client will be ignored.
	ChainRpc string

	// Client is the client that wraps the database and connects to the network.
	//
	// TODO: When ChainRpc is empty, the Client will be used. When ChainRpc is provided, a new client
	// will be created from the rpc url and Client will be ignored.
	Client client.Client

	// StartBlock is the starting block height from which the process will start when no record of the
	// process exists in the database. When a record does exist, start block is ignored.
	StartBlock int64
}

func (p Params) ValidateParams() error {
	if p.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if p.ChainId == "" {
		return fmt.Errorf("chain id cannot be empty")
	}

	if p.StartBlock == 0 {
		return fmt.Errorf("start block cannot be empty")
	}

	return nil
}
