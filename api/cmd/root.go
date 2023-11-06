package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/chora-io/server/api/app"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "api",
		Short:   "An application interface for chora server",
		Long:    "An application interface for chora server",
		Example: "api",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			// load config
			cfg := app.LoadConfig()

			// set logger
			log := zerolog.New(os.Stdout)

			// initialize application
			app := app.Initialize(cfg, log)

			// set the host address
			host := fmt.Sprintf(":%d", cfg.ApiPort)

			// run application
			app.Run(host)

			return nil
		},
	}

	return cmd
}
