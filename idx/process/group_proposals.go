package process

import (
	"context"
	"fmt"

	"github.com/choraio/server/idx/client"
)

const GroupProposalsName = "group-proposals"

func GroupProposals(ctx context.Context, c client.Client, chainId string) error {
	// get last processed block from database
	lastBlock, err := c.GetProcessLastBlock(ctx, chainId, GroupProposalsName)
	if err != nil {
		return err
	}

	fmt.Println("last block", lastBlock)

	// query block for proposal pruned events
	events, err := c.GetGroupEventProposalPruned(lastBlock)
	if err != nil {
		return err
	}

	for i, event := range events {
		fmt.Println("event", i, event)

		// fetch proposal from archive node
		proposal, err := c.GetGroupProposalAtBlockHeight(lastBlock, "1")
		if err != nil {
			return err
		}

		fmt.Println("proposal", proposal)

		// TODO: store group proposal in database
	}

	// increment last processed block in database
	err = c.UpdateProcessLastBlock(ctx, chainId, GroupProposalsName, lastBlock+1)
	if err != nil {
		return err
	}

	return nil
}
