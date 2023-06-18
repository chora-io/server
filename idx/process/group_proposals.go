package process

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/choraio/server/idx/client"
)

func GroupProposals(ctx context.Context, c client.Client, p Params) error {
	fmt.Println(p.Name, "start")

	// select last processed block from database
	lastBlock, err := c.SelectProcessLastBlock(ctx, p.ChainId, p.Name)

	// handle no rows error
	if err == sql.ErrNoRows {
		// set last block to start block
		lastBlock = p.StartBlock

		// insert process starting at start block
		err = c.InsertProcessLastBlock(ctx, p.ChainId, p.Name, p.StartBlock)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	fmt.Println(p.Name, "last block", lastBlock)

	// query block for proposal pruned events
	events, err := c.GetGroupEventProposalPruned(lastBlock)
	if err != nil {
		return err
	}

	for i, event := range events {
		fmt.Println(p.Name, "event", i, event)

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

		fmt.Println(p.Name, "proposal", proposal)

		// add group proposal to database
		err = c.InsertGroupProposal(ctx, p.ChainId, proposalId, proposal)
		if err != nil {
			return err
		}
	}

	// increment last processed block in database
	err = c.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, lastBlock+1)
	if err != nil {
		return err
	}

	fmt.Println(p.Name, "end")

	return nil
}
