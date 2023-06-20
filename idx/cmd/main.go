package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	db "github.com/choraio/server/db/client"
	"github.com/choraio/server/idx/client"
	"github.com/choraio/server/idx/config"
	"github.com/choraio/server/idx/process"
	"github.com/choraio/server/idx/runner"
	"github.com/rs/zerolog"
)

func main() {
	// set context signalling cancellation when SIGINT or SIGTERM is received
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// load config
	cfg := config.LoadConfig()

	// set logger
	log := zerolog.New(os.Stdout)

	// initialize and set db client
	db, err := db.NewDatabase(cfg.DatabaseUrl, log)
	if err != nil {
		panic(err)
	}

	// create clients that wrap db and logger
	c, err := client.NewClient("127.0.0.1:9090", db, log)
	if err != nil {
		panic(err)
	}
	// ...

	// create process runner
	r := runner.NewRunner(ctx, cfg)

	// run processes
	r.RunProcess(process.GroupProposals, process.Params{
		Name:       "group-proposals",
		ChainId:    "chora-local",
		Client:     c,
		StartBlock: 1,
	})
	// ...

	// shut down runner
	r.Close()

	// shut down db
	db.Close()

	// shut down clients
	c.Close()
	// ...
}
