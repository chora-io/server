package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	db "github.com/choraio/server/db/client"
	"github.com/choraio/server/idx/client"
	"github.com/choraio/server/idx/config"
	"github.com/choraio/server/idx/process"
	"github.com/choraio/server/idx/runner"
)

const (
	FlagStartBlock string = "start-block"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "idx [rpc] [chain-id]",
		Short:   "A process runner for indexing blockchain state",
		Long:    "A process runner for indexing blockchain state",
		Example: "idx localhost:9090 chora-local --start-block 1",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// set context signalling cancellation when SIGINT or SIGTERM is received
			ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

			// load config
			cfg := config.LoadConfig()

			// set logger
			log := zerolog.New(os.Stdout)

			// initialize and set db client
			db, err := db.NewDatabase(cfg.DatabaseUrl, log)
			if err != nil {
				return err
			}

			// create client that wraps db and logger
			c, err := client.NewClient(args[0], db, log)
			if err != nil {
				return err
			}

			// get start block (default 1)
			startBlock, err := cmd.Flags().GetUint64(FlagStartBlock)
			if err != nil {
				return err
			}

			// set parameters for each process
			groupProposalsParams := process.Params{
				Name:       "group-proposals",
				ChainId:    args[1],
				Client:     c,
				StartBlock: startBlock,
			}
			groupVotesParams := process.Params{
				Name:       "group-votes",
				ChainId:    args[1],
				Client:     c,
				StartBlock: startBlock,
			}
			// ...

			// validate process parameters for each process
			err = groupProposalsParams.ValidateParams()
			if err != nil {
				return err
			}
			err = groupVotesParams.ValidateParams()
			if err != nil {
				return err
			}
			// ...

			// create process runner
			r := runner.NewRunner(ctx, cfg)

			// run processes
			r.RunProcess(process.GroupProposals, groupProposalsParams)
			r.RunProcess(process.GroupVotes, groupVotesParams)
			// ...

			// shut down runner
			r.Close()

			// shut down db
			db.Close()

			// shut down client
			c.Close()

			return nil
		},
	}

	cmd.Flags().Uint64(FlagStartBlock, 1, "the starting block for a new process")

	return cmd
}
