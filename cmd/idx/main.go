package main

import (
	"fmt"

	"github.com/choraio/server/idx/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
	}
}
