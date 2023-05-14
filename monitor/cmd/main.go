package cmd

import (
	"fmt"

	"github.com/choraio/server/monitor/app"
)

// nolint
func main() {
	cfg := app.LoadConfig()
	fmt.Println(cfg)
}
