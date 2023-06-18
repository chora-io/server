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
	cdc  *GRPCCodec
	conn *grpc.ClientConn
}

// NewClient creates a new client.
func NewClient(rpcUrl string) (Client, error) {
	c := Client{}
	c.ctx = context.Background()
	c.cdc = CustomCodec()

	conn, err := grpc.Dial(
		rpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(c.cdc)),
	)
	if err != nil {
		return Client{}, err
	}
	c.conn = conn

	return c, nil
}

// Close closes the client.
func (c Client) Close() error {
	return nil
}

// GetGroupEventProposalPruned gets group.v1.EventProposalPruned from a given block.
func (c Client) GetGroupEventProposalPruned(height int64) ([]json.RawMessage, error) {
	block, err := c.getBlockByHeight(c.ctx, height)
	if err != nil {
		return nil, err
	}

	fmt.Println("block", block)

	for _, bz := range block.Data.GetTxs() {
		var tx types.Tx

		err := c.cdc.Unmarshal(bz, &tx)
		if err != nil {
			return nil, err
		}

		for _, msg := range tx.GetMsgs() {
			fmt.Println("msg", msg.String())
		}
	}

	submitEvents, err := c.getTxsEvent(c.ctx, "cosmos.group.v1.MsgSubmitProposal", height)
	if err != nil {
		return nil, err
	}

	fmt.Println("submit events", submitEvents)

	return nil, nil
}

// GetGroupProposal gets a group proposal by proposal id at a given block height.
func (c Client) GetGroupProposal(height int64, proposalId int64) (json.RawMessage, error) {
	resp, err := group.NewQueryClient(c.conn).Proposal(c.ctx, &group.QueryProposalRequest{
		ProposalId: uint64(proposalId),
	})
	if err != nil {
		return nil, err
	}

	bz, err := json.Marshal(resp.Proposal)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func (c Client) getBlockByHeight(ctx context.Context, height int64) (*tmservice.Block, error) {
	res, err := tmservice.NewServiceClient(c.conn).GetBlockByHeight(ctx, &tmservice.GetBlockByHeightRequest{
		Height: height,
	})
	if err != nil {
		return nil, err
	}

	return res.SdkBlock, nil
}

func (c Client) getTxsEvent(ctx context.Context, msgTypeName string, height int64) (*tx.GetTxsEventResponse, error) {
	svc := tx.NewServiceClient(c.conn)
	return svc.GetTxsEvent(ctx, &tx.GetTxsEventRequest{Events: []string{
		fmt.Sprintf(`message.action='/%s'`, msgTypeName),
		fmt.Sprintf(`tx.height=%d`, height),
	}})
}
