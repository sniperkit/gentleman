package main

import (
	"fmt"

	gentleman "github.com/sniperkit/gentleman/pkg"
	"github.com/sniperkit/gentleman/plugin/compression"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Disable HTTP compression
	cli.Use(compression.Disable())

	// Perform the request
	res, err := cli.Request().URL("http://httpbin.org/gzip").Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body: %s", res.String())
}
