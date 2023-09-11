package process

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/group"
)

// GroupProposals is a process function for indexing group proposals.
func GroupProposals(ctx context.Context, p Params) error {

	// determine last block and next block for process
	lastBlock, nextBlock, err := AdvanceProcess(ctx, p)
	if err != nil {
		return err
	} else if lastBlock == nextBlock {
		return nil // do nothing because process is in sync
	}

	// query next block for proposal pruned events
	events, err := p.Client.GetGroupEventProposalPruned(nextBlock)
	if err != nil {
		return err
	}

	for _, event := range events {
		// get proposal id from event
		proposalId := event.ProposalId

		// fetch proposal at last block height
		proposal, groupId, err := p.Client.GetGroupProposal(lastBlock, proposalId)
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

		// unmarshal proposal so that we can check and update
		var update group.Proposal
		err = p.Client.Codec().UnmarshalJSON(proposal, &update)
		if err != nil {
			return err
		}

		var updated json.RawMessage

		// TODO: handle proposal accepted but not executed..?
		if update.Status == 2 {

			fmt.Println(p.Name, "updating group proposal executor result", p.ChainId, proposalId)

			// update executor result from not run to success
			update.ExecutorResult = group.ProposalExecutorResult(2)

			// marshal updated proposal
			updated, err = p.Client.Codec().MarshalJSON(&update)
			if err != nil {
				return err
			}
		}

		fmt.Println(p.Name, "adding group proposal", p.ChainId, proposalId)

		// add group proposal to database
		err = p.Client.InsertGroupProposal(ctx, p.ChainId, proposalId, groupId, updated)
		if err != nil && strings.Contains(err.Error(), "duplicate key value") {
			fmt.Println(p.Name, "group proposal exists", p.ChainId, proposalId)

			fmt.Println(p.Name, "updating group proposal", p.ChainId, proposalId)

			// update group proposal in database
			err = p.Client.UpdateGroupProposal(ctx, p.ChainId, proposalId, updated)
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
