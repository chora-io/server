package cosmos

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"github.com/cosmos/cosmos-sdk/types/tx"
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
	conn, err := grpc.Dial(
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

// GetGroupEventProposalPruned gets any array of group.v1.EventProposalPruned from block height.
func (c Client) GetGroupEventProposalPruned(height int64) ([]group.EventProposalPruned, error) {
	// get all transactions from block height
	txs, err := tx.NewServiceClient(c.conn).GetTxsEvent(c.ctx, &tx.GetTxsEventRequest{
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
			if e.Type == "cosmos.group.v1.EventProposalPruned" {

				protoEvent, err := types.ParseTypedEvent(e)
				if err != nil {
					return nil, err
				}

				event, ok := protoEvent.(*group.EventProposalPruned)
				if !ok {
					return nil, fmt.Errorf("expected %T got %T", group.EventProposalPruned{}, protoEvent)
				}

				events = append(events, *event)
			}
		}
	}

	return events, nil
}

// GetGroupProposal gets a group proposal by proposal id at block height.
func (c Client) GetGroupProposal(height int64, proposalId int64) (json.RawMessage, error) {
	// convert block height to string
	blockHeight := strconv.FormatInt(height, 10)

	// set context to use block height in header
	ctx := metadata.AppendToOutgoingContext(c.ctx, grpctypes.GRPCBlockHeightHeader, blockHeight)

	// query proposal at block height using context with block height
	resp, err := group.NewQueryClient(c.conn).Proposal(ctx, &group.QueryProposalRequest{
		ProposalId: uint64(proposalId),
	})
	if err != nil {
		return nil, err
	}

	// TODO: fix encoding of nested any types
	// go run ./cmd/idx testnet.chora.io:9090 chora-testnet-1 2646260
	for i, m := range resp.Proposal.Messages {
		if m.TypeUrl == "/cosmos.group.v1.MsgUpdateGroupPolicyDecisionPolicy" {
			resp.Proposal.Messages[i] = nil
		}
	}

	// get json encoding of proposal
	bz, err := json.Marshal(resp.Proposal)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

// GetLatestBlockHeight gets the latest block height.
func (c Client) GetLatestBlockHeight() (int64, error) {
	res, err := tmservice.NewServiceClient(c.conn).GetLatestBlock(c.ctx, &tmservice.GetLatestBlockRequest{})
	if err != nil {
		return 0, err
	}
	return res.SdkBlock.Header.Height, nil
}
