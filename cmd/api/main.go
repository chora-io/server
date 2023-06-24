package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/choraio/server/api/app"
)

func main() {
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
}
