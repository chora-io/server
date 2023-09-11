package process

import (
	"context"
	"fmt"
	"strings"
)

// GroupVotes is a process function for indexing group votes.
func GroupVotes(ctx context.Context, p Params) error {

	// determine last block and next block for process
	lastBlock, nextBlock, err := AdvanceProcess(ctx, p)
	if err != nil {
		return err
	} else if lastBlock == nextBlock {
		return nil // do nothing because process is in sync
	}

	// query next block for vote events
	events, err := p.Client.GetGroupEventVote(nextBlock)
	if err != nil {

		// TODO: resolve codec errors and reconsider skipped blocks

		fmt.Println(p.Name, "error", p.ChainId, nextBlock, err.Error())

		// insert skipped block and error message to retry in a separate process
		err := p.Client.InsertSkippedBlock(ctx, p.ChainId, p.Name, nextBlock, err.Error())
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				// skipped block already recorded
			} else {
				return err
			}
		}
	}

	for _, event := range events {
		// get proposal id from event
		proposalId := event.ProposalId

		// get voter address from event
		voter := event.Voter

		// fetch vote at next block height
		vote, err := p.Client.GetGroupVote(nextBlock, proposalId, voter)
		if err != nil {

			// TODO: resolve codec errors and reconsider skipped blocks

			fmt.Println(p.Name, "error", p.ChainId, nextBlock, err.Error())

			// insert skipped block and error message to retry in a separate process
			err := p.Client.InsertSkippedBlock(ctx, p.ChainId, p.Name, nextBlock, err.Error())
			if err != nil {
				if strings.Contains(err.Error(), "duplicate key value") {
					continue // skipped block already recorded
				}
				return err
			}

			// skip this event
			continue
		}

		fmt.Println(p.Name, "adding group vote", p.ChainId, proposalId, voter)

		// add group vote to database
		err = p.Client.InsertGroupVote(ctx, p.ChainId, proposalId, voter, vote)
		if err != nil && strings.Contains(err.Error(), "duplicate key value") {
			fmt.Println(p.Name, "group vote exists", p.ChainId, proposalId, voter)

			fmt.Println(p.Name, "updating group vote", p.ChainId, proposalId, voter)

			// update group vote in database
			err = p.Client.UpdateGroupVote(ctx, p.ChainId, proposalId, voter, vote)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	// increment last processed block in database
	err = p.Client.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, nextBlock)
	if err != nil {
		return err
	}

	return nil
}
