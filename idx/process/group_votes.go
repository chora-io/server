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
