package cosmos

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/group"
)

// Client is the client.
type Client struct {
	ctx  context.Context
	cdc  *Codec
	conn *grpc.ClientConn
}

// NewClient creates a new client.
func NewClient(rpcUrl string) (Client, error) {
	c := Client{}

	// set context
	c.ctx = context.Background()

	// set custom codec
	c.cdc = CustomCodec()

	// make rpc connection
	conn, err := grpc.NewClient(
		rpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(c.cdc)),
	)
	if err != nil {
		return Client{}, err
	}

	// set connection
	c.conn = conn

	return c, nil
}

// Close shuts down the client.
func (c Client) Close() error {

	// close the connection
	err := c.conn.Close()
	if err != nil {
		return err
	}

	return nil
}

// Codec returns the client codec.
func (c Client) Codec() *Codec {
	return c.cdc
}

// GetGroupEventProposalPruned gets an array of group.v1.EventProposalPruned from block height.
func (c Client) GetGroupEventProposalPruned(height int64) ([]group.EventProposalPruned, error) {

	// get all transactions from block height
	txs, err := sdktx.NewServiceClient(c.conn).GetTxsEvent(c.ctx, &sdktx.GetTxsEventRequest{
		Events: []string{
			fmt.Sprintf(`tx.height=%d`, height),
		},
	})
	if err != nil {
		return nil, err
	}

	// initialize event proposal pruned array
	events := make([]group.EventProposalPruned, 0)

	// loop through transaction responses
	for _, tx := range txs.TxResponses {

		// ignore failed transactions
		if tx.Code != 0 {
			continue
		}

		// loop through transaction events
		for _, e := range tx.Events {

			// parse and append proposal pruned events
			if e.Type == "cosmos.group.v1.EventProposalPruned" {

				// parse typed event
				protoEvent, err := types.ParseTypedEvent(e)
				if err != nil {
					return nil, err
				}

				// type cast parsed event
				event, ok := protoEvent.(*group.EventProposalPruned)
				if !ok {
					return nil, fmt.Errorf("expected %T got %T", group.EventProposalPruned{}, protoEvent)
				}

				// append type cast event
				events = append(events, *event)
			}
		}
	}

	return events, nil
}

// GetGroupEventVote gets an array of group.v1.EventVote from block height. This method
// also returns the voter address pulled from the tx message so that the vote can be queried
// by proposal id and voter address (i.e. voter address is not provided by EventVote).
func (c Client) GetGroupEventVote(height uint64) ([]EventVoteWithVoter, error) {

	// get all transactions from block height
	txs, err := sdktx.NewServiceClient(c.conn).GetTxsEvent(c.ctx, &sdktx.GetTxsEventRequest{
		Events: []string{
			fmt.Sprintf(`tx.height=%d`, height),
		},
	})
	if err != nil {
		return nil, err
	}

	// initialize event vote array
	events := make([]EventVoteWithVoter, 0)

	// loop through transaction responses
	for _, txr := range txs.TxResponses {

		// ignore failed transactions
		if txr.Code != 0 {
			continue
		}

		// declare voter (for voter workaround)
		var voter string

		// unmarshal transaction (for voter workaround)
		var tx sdktx.Tx
		err := c.cdc.Unmarshal(txr.Tx.Value, &tx)
		if err != nil {

			// TODO: handle unregistered messages or register all to codec

			return nil, err
		}

		// loop through transaction messages (for voter workaround)
		for _, m := range tx.Body.Messages {

			// TODO: handle MsgVote wrapped in anther message (gov proposal, group proposal, authz exec, intertx)

			// NOTE: If there are two MsgVote messages within the same transaction,
			// the transaction will fail because only one can be executed successfully
			// given there is only one signer per transaction, therefore, we are not
			// concerned if there are multiple MsgVote within the same transaction.

			// find first vote message within transaction messages
			if voter == "" && m.TypeUrl == "/cosmos.group.v1.MsgVote" {

				// unmarshal vote message
				var msgVote group.MsgVote
				err := c.cdc.Unmarshal(m.Value, &msgVote)
				if err != nil {
					return nil, err
				}

				// set voter address
				voter = msgVote.Voter
			}
		}

		// loop through transaction events
		for _, e := range txr.Events {

			// parse and append vote events
			if e.Type == "cosmos.group.v1.EventVote" {

				// parse typed event
				protoEvent, err := types.ParseTypedEvent(e)
				if err != nil {
					return nil, err
				}

				// type cast parsed event
				event, ok := protoEvent.(*group.EventVote)
				if !ok {
					return nil, fmt.Errorf("expected %T got %T", group.EventVote{}, protoEvent)
				}

				// append type cast event (including voter)
				events = append(events, EventVoteWithVoter{
					ProposalId: event.ProposalId,
					Voter:      voter,
				})
			}
		}
	}

	return events, nil
}

// GetGroupProposal gets a group proposal by proposal id at block height.
func (c Client) GetGroupProposal(height uint64, proposalId uint64) (json.RawMessage, uint64, error) {

	// convert block height to string
	blockHeight := strconv.FormatUint(height, 10)

	// set context to use block height in header
	ctx := metadata.AppendToOutgoingContext(c.ctx, grpctypes.GRPCBlockHeightHeader, blockHeight)

	// query proposal at block height using context with block height
	resp, err := group.NewQueryClient(c.conn).Proposal(ctx, &group.QueryProposalRequest{
		ProposalId: proposalId,
	})
	if err != nil {
		return nil, 0, err
	}

	// query policy at block height using context with block height
	policyResp, err := group.NewQueryClient(c.conn).GroupPolicyInfo(ctx, &group.QueryGroupPolicyInfoRequest{
		Address: resp.Proposal.GroupPolicyAddress,
	})
	if err != nil {
		return nil, 0, err
	}

	// get json encoding of proposal
	bz, err := c.cdc.MarshalJSON(resp.Proposal)
	if err != nil {
		return nil, 0, err
	}

	return bz, policyResp.Info.GroupId, nil
}

// GetGroupVote gets a group vote by proposal id and voter address.
func (c Client) GetGroupVote(height uint64, proposalId uint64, voter string) (json.RawMessage, error) {

	// convert block height to string
	blockHeight := strconv.FormatUint(height, 10)

	// set context to use block height in header
	ctx := metadata.AppendToOutgoingContext(c.ctx, grpctypes.GRPCBlockHeightHeader, blockHeight)

	// query vote at block height using context with block height
	resp, err := group.NewQueryClient(c.conn).VoteByProposalVoter(ctx, &group.QueryVoteByProposalVoterRequest{
		ProposalId: proposalId,
		Voter:      voter,
	})
	if err != nil {
		return nil, err
	}

	// get json encoding of vote
	bz, err := c.cdc.MarshalJSON(resp.Vote)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

// GetLatestBlockHeight gets the latest block height.
func (c Client) GetLatestBlockHeight() (uint64, error) {

	// get latest block
	res, err := cmtservice.NewServiceClient(c.conn).GetLatestBlock(c.ctx, &cmtservice.GetLatestBlockRequest{})
	if err != nil {
		return 0, err
	}

	// return latest block height
	return uint64(res.SdkBlock.Header.Height), nil
}
