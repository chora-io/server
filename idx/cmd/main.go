package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/choraio/server/idx/config"
	context2 "github.com/choraio/server/idx/context"
	"github.com/choraio/server/idx/process"
	"github.com/choraio/server/idx/runner"
)

func main() {
	// set context signalling cancellation when SIGINT or SIGTERM is received
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// load configuration
	cfg := config.LoadConfig()

	// create indexer context
	idxCtx, err := context2.NewContext(cfg)
	if err != nil {
		panic(err)
	}

	// create runner
	r := runner.NewRunner(ctx, idxCtx)

	// run processes
	r.RunProcess("example-1", process.Example)
	r.RunProcess("example-2", process.Example)

	r.Close()
}
