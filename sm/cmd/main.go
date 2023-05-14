package cmd

import (
	"fmt"
	"os"

	"github.com/choraio/server/sm/client"
)

// nolint
func main() {
	c := client.NewClient("http://localhost:3030/resources")

	bz, err := os.ReadFile("../data/example.ttl")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	res, err := c.Post(bz)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Println("post response:", res)

	res, err = c.Get()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Println("get response:", res)
}
