package cmd

import (
	"fmt"

	"github.com/choraio/server/idx/app"
)

// nolint
func main() {
	cfg := app.LoadConfig()
	fmt.Println(cfg)
}
