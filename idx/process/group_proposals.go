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
		proposalId := int64(event.ProposalId)

		// fetch proposal at last block height
		proposal, groupId, err := p.Client.GetGroupProposal(lastBlock, proposalId)

		// TODO: handle proposal not found error
		if err != nil {
			return err
		}

		// unmarshal proposal so that we can check and update
		var update group.Proposal
		err = json.Unmarshal(proposal, &update)
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
			updated, err = json.Marshal(update)
			if err != nil {
				return err
			}
		}

		fmt.Println(p.Name, "adding group proposal", p.ChainId, proposalId)

		// add group proposal to database
		err = p.Client.InsertGroupProposal(ctx, p.ChainId, proposalId, groupId, updated)
		if err != nil && strings.Contains(err.Error(), "duplicate key value ") {
			fmt.Println(p.Name, "error", err.Error())

			fmt.Println(p.Name, "updating group proposal", p.ChainId, proposalId)

			// update group proposal in database
			err = p.Client.UpdateGroupProposal(ctx, p.ChainId, proposalId, proposal)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		fmt.Println(p.Name, "successfully processed event", p.ChainId, event.String())
		fmt.Println(p.Name, "successfully added proposal", p.ChainId, proposalId)
	}

	// increment last processed block in database
	err = p.Client.UpdateProcessLastBlock(ctx, p.ChainId, p.Name, nextBlock)
	if err != nil {
		return err
	}

	return nil
}
