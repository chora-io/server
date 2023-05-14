package cmd

import (
	"fmt"

	"github.com/choraio/server/indexer/app"
)

// nolint
func main() {
	cfg := app.LoadConfig()
	fmt.Println(cfg)
}
