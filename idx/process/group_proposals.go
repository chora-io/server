package process

import (
	"context"
	"fmt"
	"strconv"

	"github.com/choraio/server/idx/client"
)

func GroupProposals(ctx context.Context, c client.Client, p Params) error {
	// get last processed block from database
	lastBlock, err := c.GetProcessLastBlock(ctx, p.ChainId, p.ProcessName)
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

		// TODO: get proposal id from event
		proposalId, err := strconv.ParseInt("1", 0, 64)
		if err != nil {
			return err
		}

		// fetch proposal at last block height
		proposal, err := c.GetGroupProposal(lastBlock, proposalId)
		if err != nil {
			return err
		}

		fmt.Println("proposal", proposal)

		// add group proposal to database
		err = c.AddGroupProposal(ctx, p.ChainId, proposalId, proposal)
		if err != nil {
			return err
		}
	}

	// increment last processed block in database
	err = c.UpdateProcessLastBlock(ctx, p.ChainId, p.ProcessName, lastBlock+1)
	if err != nil {
		return err
	}

	return nil
}
