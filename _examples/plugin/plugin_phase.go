package main

import (
	"fmt"

	gentleman "github.com/sniperkit/gentleman/pkg"
	"github.com/sniperkit/gentleman/pkg/context"
	"github.com/sniperkit/gentleman/pkg/plugin"

	"github.com/sniperkit/gentleman/plugin/headers"
)

func main() {
	// Create a new client
	cli := gentleman.New()

	// Define a custom header
	cli.Use(headers.Set("Token", "s3cr3t"))

	// Create a plugin for the response phase
	cli.Use(plugin.NewPhasePlugin("response", func(ctx *context.Context, h context.Handler) {
		ctx.Response.StatusCode = 201 // change the status code
		h.Next(ctx)
	}))

	// Perform the request
	res, err := cli.Request().URL("http://httpbin.org/headers").Send()
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
