package main

import (
	"fmt"

	"github.com/choraio/server/iri/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
	}
}
