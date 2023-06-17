package process

import (
	"context"
	"fmt"

	"github.com/choraio/server/idx/client"
)

const GroupProposalsName = "group-proposals"

func GroupProposals(ctx context.Context, c client.Client, chainId string) error {
	// get last processed block from database
	lastBlock, err := c.GetLastBlock(ctx, chainId, GroupProposalsName)
	if err != nil {
		return err
	}

	fmt.Println("last block", lastBlock)

	// TODO: query block for group events

	// TODO: for each group event...
	// - fetch proposal from archive node
	// - store group proposal in database

	// increment last processed block in database
	err = c.IncrementLastBlock(ctx, chainId, GroupProposalsName)
	if err != nil {
		return err
	}

	return nil
}
