package process

// Params are the process parameters.
type Params struct {
	// Name is the name of the process.
	Name string

	// ChainId is the chain id of the network (e.g. chora-testnet-1, regen-redwood-1).
	ChainId string

	// StartBlock is the starting block height from which the process will start when no record of the
	// process exists in the database. When a record does exist, start block is ignored.
	StartBlock int64
}
