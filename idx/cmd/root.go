package cmd

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	db "github.com/choraio/server/db/client"
	"github.com/choraio/server/idx/client"
	"github.com/choraio/server/idx/config"
	"github.com/choraio/server/idx/process"
	"github.com/choraio/server/idx/runner"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "idx [rpc] [chain-id] [start-block]",
		Short:   "A process runner for indexing blockchain state",
		Long:    "A process runner for indexing blockchain state",
		Example: "idx localhost:9090 chora-local 1",
		Args:    cobra.ExactArgs(3),
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

			// create clients that wrap db and logger
			c, err := client.NewClient(args[0], db, log)
			if err != nil {
				return err
			}
			// ...

			// parse start block
			startBlock, err := strconv.ParseInt(args[2], 0, 64)
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
			// ...

			// validate process parameters for each process
			err = groupProposalsParams.ValidateParams()
			if err != nil {
				return err
			}
			// ...

			// create process runner
			r := runner.NewRunner(ctx, cfg)

			// run processes
			r.RunProcess(process.GroupProposals, groupProposalsParams)
			// ...

			// shut down runner
			r.Close()

			// shut down db
			db.Close()

			// shut down clients
			c.Close()
			// ...

			return nil
		},
	}

	return cmd
}
