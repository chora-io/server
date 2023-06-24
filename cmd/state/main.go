package main

import (
	"fmt"
	"os"

	"github.com/choraio/server/state/app"
)

func main() {
	app := app.NewApp("http://localhost:3030/resources")

	bz, err := os.ReadFile("../data/example.ttl")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	res, err := app.Post(bz)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Println("post response:", res)

	res, err = app.Get()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Println("get response:", res)
}
