package process

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/choraio/server/idx/client"
)

func GroupProposals(ctx context.Context, c client.Client, p Params) error {
	// get latest block height from chain
	latestBlock, err := c.GetLatestBlockHeight()
	if err != nil {
		return err
	}

	// select last processed block from database
	lastBlock, err := c.SelectProcessLastBlock(ctx, p.ChainId, p.Name)

	// handle process in sync
	if lastBlock == latestBlock {
		fmt.Println(p.Name, "process is in sync with latest block")

		fmt.Println(p.Name, "last block", lastBlock)
		fmt.Println(p.Name, "latest block", latestBlock)

		return nil // do nothing because process is in sync with latest block
	}

	// handle process mismatch
	if latestBlock < lastBlock {
		fmt.Println(p.Name, "updating last processed block to latest block")

		// set last block to latest block
		lastBlock = latestBlock

		// update last processed block to latest block
		err = c.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, latestBlock)
		if err != nil {
			return err
		}
	}

	// handle last block error
	if err == sql.ErrNoRows {
		fmt.Println(p.Name, "inserting last processed block as start block")

		// set last block to start block
		lastBlock = p.StartBlock

		// insert last processed block as start block
		err = c.InsertProcessLastBlock(ctx, p.ChainId, p.Name, p.StartBlock)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	fmt.Println(p.Name, "last block", lastBlock)

	nextBlock := lastBlock + 1

	fmt.Println(p.Name, "next block", nextBlock)

	// query block for proposal pruned events
	events, err := c.GetGroupEventProposalPruned(nextBlock)
	if err != nil {
		return err
	}

	for _, event := range events {
		// get proposal id from event
		proposalId := int64(event.ProposalId)

		// fetch proposal at last block height
		proposal, err := c.GetGroupProposal(nextBlock, proposalId)

		// handle proposal not found error
		if err != nil {
			fmt.Println(p.Name, "proposal not found", proposalId)

			return err // TODO: alert and continue
		}

		// add group proposal to database
		err = c.InsertGroupProposal(ctx, p.ChainId, proposalId, proposal)
		if err != nil {
			return err
		}
	}

	// increment last processed block in database
	err = c.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, nextBlock)
	if err != nil {
		return err
	}

	return nil
}
