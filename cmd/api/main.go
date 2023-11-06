package main

import (
	"fmt"

	"github.com/chora-io/server/api/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
	}
}
