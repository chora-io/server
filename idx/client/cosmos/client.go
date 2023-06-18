package cosmos

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

// GetGroupEventProposalPruned gets group.v1.EventProposalPruned from a given block.
func (c Client) GetGroupEventProposalPruned(height int64) ([]json.RawMessage, error) {
	txs, err := c.getTxsEvent(c.ctx, "cosmos.group.v1.EventProposalPruned", height)
	if err != nil {
		return nil, err
	}

	fmt.Println("transactions with event", txs.Total)

	// TODO: get events from transactions

	return nil, nil
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

func (c Client) getTxsEvent(ctx context.Context, msgTypeName string, height int64) (*tx.GetTxsEventResponse, error) {
	svc := tx.NewServiceClient(c.conn)
	return svc.GetTxsEvent(ctx, &tx.GetTxsEventRequest{Events: []string{
		fmt.Sprintf(`message.action='/%s'`, msgTypeName),
		fmt.Sprintf(`tx.height=%d`, height),
	}})
}
