package client

// CosmosClient is the client.
type CosmosClient struct {
	chainId string
	rpcUrl  string
}

// NewCosmosClient creates a new client.
func NewCosmosClient(chainId, rpcUrl string) (CosmosClient, error) {
	c := CosmosClient{}
	c.chainId = chainId
	c.rpcUrl = rpcUrl
	return c, nil
}

// Close closes the client.
func (c CosmosClient) Close() error {
	return nil
}

// GetGroupEventProposalPruned gets group.v1.EventProposalPruned from a given block.
func (c CosmosClient) GetGroupEventProposalPruned(block int64) ([]any, error) {
	return nil, nil
}

// GetGroupProposalAtBlockHeight gets a group proposal by proposal id at a given block height.
func (c CosmosClient) GetGroupProposalAtBlockHeight(block int64, proposalId string) (any, error) {
	return nil, nil
}
