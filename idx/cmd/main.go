package main

import (
	"context"
	"fmt"
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

var rootCmd = &cobra.Command{
	Use:     "idx [db] [rpc] [chain-id] [start-block]",
	Short:   "A process runner for indexing blockchain state",
	Long:    "A process runner for indexing blockchain state",
	Example: "idx postgres://postgres:password@localhost:5432/postgres?sslmode=disable localhost:9090 chora-local 1",
	Args:    cobra.ExactArgs(4),
	RunE: func(cmd *cobra.Command, args []string) error {
		// set context signalling cancellation when SIGINT or SIGTERM is received
		ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

		// load config
		cfg := config.LoadConfig()

		// set logger
		log := zerolog.New(os.Stdout)

		// initialize and set db client
		db, err := db.NewDatabase(args[0], log)
		if err != nil {
			return err
		}

		// creates a client that wraps db and logger
		c, err := client.NewClient(args[1], db, log)
		if err != nil {
			return err
		}
		// ...

		// create process runner
		r := runner.NewRunner(ctx, cfg)

		// parse start block
		startBlock, err := strconv.ParseInt(args[3], 0, 64)
		if err != nil {
			return err
		}

		// run processes
		r.RunProcess(process.GroupProposals, process.Params{
			Name:       "group-proposals",
			ChainId:    args[2],
			Client:     c,
			StartBlock: startBlock,
		})
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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
