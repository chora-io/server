package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/choraio/server/app"
	"github.com/choraio/server/db"
)

func main() {
	cfg := app.LoadConfig()
	log := zerolog.New(os.Stdout)
	db, err := db.NewDatabase(cfg.AppDatabaseUrl, log)
	if err != nil {
		panic(err)
	}
	app := app.Initialize(cfg, db.Reader(), db.Writer(), log)
	app.Run(fmt.Sprintf(":%d", cfg.AppPort))
}
