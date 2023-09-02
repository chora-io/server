package process

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func GroupVotes(ctx context.Context, p Params) error {
	// get latest block from configured client
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

		// set last processed block to latest block
		lastBlock = latestBlock

		// update last processed block to latest block
		err = p.Client.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, latestBlock)
		if err != nil {
			return err
		}
	}

	// handle process error
	if err == sql.ErrNoRows {
		fmt.Println(p.Name, "inserting (start block - 1) as last processed block")

		// set last processed block to (start block - 1)
		lastBlock = p.StartBlock - 1

		// insert (start block - 1) as last processed block
		err = p.Client.InsertProcessLastBlock(ctx, p.ChainId, p.Name, lastBlock)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	fmt.Println(p.Name, "last block", lastBlock)

	nextBlock := lastBlock + 1

	fmt.Println(p.Name, "next block", nextBlock)

	// query next block for vote events
	events, err := p.Client.GetGroupEventVote(nextBlock)
	if err != nil {
		return err
	}

	for _, event := range events {
		// get proposal id from event
		proposalId := int64(event.ProposalId)

		// get voter address from event
		voter := event.Voter

		// fetch vote at next block height
		vote, err := p.Client.GetGroupVote(nextBlock, proposalId, voter)

		// TODO: handle vote not found error
		if err != nil {
			return err
		}

		fmt.Println(p.Name, "adding group vote", p.ChainId, proposalId, voter)

		// add group vote to database
		err = p.Client.InsertGroupVote(ctx, p.ChainId, proposalId, voter, vote)
		if err != nil && strings.Contains(err.Error(), "duplicate key value ") {
			fmt.Println(p.Name, "error", err.Error())

			fmt.Println(p.Name, "updating group vote", p.ChainId, proposalId, voter)

			// update group vote in database
			err = p.Client.UpdateGroupVote(ctx, p.ChainId, proposalId, voter, vote)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		fmt.Println(p.Name, "successfully processed event", p.ChainId, event)
		fmt.Println(p.Name, "successfully added vote", p.ChainId, proposalId, voter)
	}

	// increment last processed block in database
	err = p.Client.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, nextBlock)
	if err != nil {
		return err
	}

	return nil
}
