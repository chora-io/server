package cosmos

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types"
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

// GetGroupEventProposalPruned gets group.v1.EventProposalPruned from a given block height.
func (c Client) GetGroupEventProposalPruned(height int64) ([]group.EventProposalPruned, error) {
	// get all transactions from block height (i.e. the event is not triggered by any one message)
	txs, err := tx.NewServiceClient(c.conn).GetTxsEvent(c.ctx, &tx.GetTxsEventRequest{
		Events: []string{
			fmt.Sprintf(`tx.height=%d`, height),
		},
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("total transactions", txs.Total)

	// initialize event proposal pruned array
	events := make([]group.EventProposalPruned, 0)

	// loop through transaction responses
	for i, tx := range txs.TxResponses {

		fmt.Println("transaction", i, tx.TxHash)

		// ignore failed transactions
		if tx.Code != 0 {
			fmt.Println("ignoring failed transaction", i, tx.TxHash)
			continue
		}

		// loop through transaction events
		for i, e := range tx.Events {

			fmt.Println("event", i, e.Type)

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

// GetGroupProposal gets a group proposal by proposal id at a given block height.
func (c Client) GetGroupProposal(height int64, proposalId int64) (json.RawMessage, error) {

	// TODO: query proposal at given block height

	resp, err := group.NewQueryClient(c.conn).Proposal(c.ctx, &group.QueryProposalRequest{
		ProposalId: uint64(proposalId),
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("response", resp)

	bz, err := json.Marshal(resp.Proposal)
	if err != nil {
		return nil, err
	}

	fmt.Println("proposal", json.RawMessage(bz))

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
