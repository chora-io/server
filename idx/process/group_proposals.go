package process

import (
	"context"
	"database/sql"
	"fmt"
)

func GroupProposals(ctx context.Context, p Params) error {
	// get latest block height from chain
	latestBlock, err := p.Client.GetLatestBlockHeight()
	if err != nil {
		return err
	}

	// select last processed block from database
	lastBlock, err := p.Client.SelectProcessLastBlock(ctx, p.ChainId, p.Name)

	// handle process in sync
	if lastBlock == latestBlock {
		fmt.Println(p.Name, "synced", lastBlock, latestBlock)
		return nil // do nothing because process is in sync
	}

	// handle process mismatch
	if latestBlock < lastBlock {
		fmt.Println(p.Name, "updating last processed block to latest block")

		// set last block to latest block
		lastBlock = latestBlock

		// update last processed block to latest block
		err = p.Client.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, latestBlock)
		if err != nil {
			return err
		}
	}

	// handle last block error
	if err == sql.ErrNoRows {
		fmt.Println(p.Name, "inserting start block as last processed block")

		// set last block to start block
		lastBlock = p.StartBlock

		// insert start block as last processed block
		err = p.Client.InsertProcessLastBlock(ctx, p.ChainId, p.Name, p.StartBlock)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	fmt.Println(p.Name, "last block", lastBlock)

	nextBlock := lastBlock + 1

	fmt.Println(p.Name, "next block", nextBlock)

	// TODO: refactor above code into reusable client method

	// query next block for proposal pruned events
	events, err := p.Client.GetGroupEventProposalPruned(nextBlock)
	if err != nil {
		return err
	}

	for _, event := range events {
		// get proposal id from event
		proposalId := int64(event.ProposalId)

		// fetch proposal at last block height
		proposal, err := p.Client.GetGroupProposal(lastBlock, proposalId)

		// handle proposal not found error
		if err != nil {
			fmt.Println(p.Name, "proposal not found", proposalId)

			return err // TODO: send alert and continue?
		}

		fmt.Println(p.Name, "inserting group proposal", p.ChainId, proposalId)

		// add group proposal to database
		err = p.Client.InsertGroupProposal(ctx, p.ChainId, proposalId, proposal)
		if err != nil {

			// TODO: pq: duplicate key value violates unique constraint "idx_group_proposal_pkey"
			fmt.Println(p.Name, "error", err.Error())

			fmt.Println(p.Name, "updating group proposal", p.ChainId, proposalId)

			// update group proposal in database
			err = p.Client.UpdateGroupProposal(ctx, p.ChainId, proposalId, proposal)
			if err != nil {
				return err
			}
		}

		fmt.Println(p.Name, "successfully processed event", p.ChainId, event.String())
		fmt.Println(p.Name, "successfully added proposal", p.ChainId, proposalId)
	}

	fmt.Println(p.Name, "updating last processed block", p.ChainId, lastBlock, nextBlock)

	// increment last processed block in database
	err = p.Client.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, nextBlock)
	if err != nil {
		return err
	}

	return nil
}
