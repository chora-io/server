package main

import (
	"fmt"

	"github.com/choraio/server/monitor/app"
)

func main() {
	cfg := app.LoadConfig()
	fmt.Println(cfg)
}
