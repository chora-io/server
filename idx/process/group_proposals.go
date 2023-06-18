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

		// TODO: event proposal
		eventProposalId := "1"

		// fetch proposal from archive node
		proposal, err := c.GetGroupProposal(lastBlock, eventProposalId)
		if err != nil {
			return err
		}

		fmt.Println("proposal", proposal)

		// parse group proposal id
		proposalId, err := strconv.ParseInt(eventProposalId, 0, 64)
		if err != nil {
			return err
		}

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
