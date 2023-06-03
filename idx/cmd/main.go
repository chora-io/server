package main

import (
	"fmt"

	"github.com/choraio/server/idx/app"
)

func main() {
	cfg := app.LoadConfig()
	fmt.Println(cfg)
}
