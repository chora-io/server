package main

import (
	"fmt"

	"github.com/choraio/server/indexer/app"
)

func main() {
	cfg := app.LoadConfig()
	fmt.Println(cfg)
}
