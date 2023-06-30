package main

import (
	"fmt"
	"os"

	"github.com/choraio/server/rdf/client"
)

func main() {
	c := client.NewClient("http://localhost:3030/resources")

	bz, err := os.ReadFile("./data/example.ttl")
	if err != nil {
		fmt.Println("error", err)
	}

	res, err := c.Post(bz)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("post response:", res)

	res, err = c.Get()
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("get response:", res)
}
