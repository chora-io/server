package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/choraio/server/idx/client"
	"github.com/choraio/server/idx/config"
	"github.com/choraio/server/idx/process"
	"github.com/choraio/server/idx/runner"
)

func main() {
	// set context signalling cancellation when SIGINT or SIGTERM is received
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// load config
	cfg := config.LoadConfig()

	// create client
	c, err := client.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// create process runner
	r := runner.NewRunner(ctx, cfg, c)

	// run processes
	r.RunProcess("group-proposals", process.GroupProposals)
	// ...

	// close process runner
	r.Close()
}