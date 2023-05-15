package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/choraio/server/api/app"
	"github.com/choraio/server/db"
)

// nolint
func main() {
	cfg := app.LoadConfig()
	log := zerolog.New(os.Stdout)
	db, err := db.NewDatabase(cfg.DatabaseUrl, log)
	if err != nil {
		panic(err)
	}
	app := app.Initialize(cfg, db.Reader(), db.Writer(), log)
	app.Run(fmt.Sprintf(":%d", cfg.ApiPort))
}
