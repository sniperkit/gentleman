package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman-mock.v2"
	"gopkg.in/h2non/gentleman.v2"
)

func main() {
	defer mock.Disable()

	// Configure the mock via gock
	mock.New("http://httpbin.org").Get("/*").Reply(204).SetHeader("Server", "gock")

	// Create a new client
	cli := gentleman.New()

	// Register the mock plugin at client level
	cli.Use(mock.Plugin)

	// Create a new request based on the current client
	req := cli.Request()

	// Define base URL
	req.URL("http://httpbin.org")

	// Define the URL path at request level
	req.Path("/status/503")

	// Set a new header field
	req.SetHeader("Client", "gentleman")

	// Perform the request
	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Header: %s\n", res.Header.Get("Server"))
}
