# [gentleman](https://github.com/h2non/gentleman)-mock [![Build Status](https://travis-ci.org/h2non/gentleman.png)](https://travis-ci.org/h2non/gentleman-mock) [![GoDoc](https://godoc.org/github.com/h2non/gentleman-mock?status.svg)](https://godoc.org/github.com/h2non/gentleman-mock) [![Coverage Status](https://coveralls.io/repos/github/h2non/gentleman-mock/badge.svg?branch=master)](https://coveralls.io/github/h2non/gentleman-mock?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/h2non/gentleman-mock)](https://goreportcard.com/report/github.com/h2non/gentleman-mock)

[gentleman](https://github.com/h2non/gentleman)'s plugin for simple HTTP mocking via [gock](https://github.com/h2non/gock).

## Installation

```bash
go get -u gopkg.in/h2non/gentleman-mock.v2
```

## Versions

- **[v1](https://github.com/h2non/gentleman-mock/tree/v1)** - First version, uses `gentleman@v1` and `gock@v1`.
- **[v2](https://github.com/h2non/gentleman-mock/tree/master)** - Latest version, uses `gentleman@v2` and `gock@v1`.

## API

See [godoc reference](https://godoc.org/github.com/h2non/gentleman-mock) for detailed API documentation.

## Example

```go
package main

import (
  "fmt"

  "gopkg.in/h2non/gentleman.v2"
  "gopkg.in/h2non/gentleman-mock.v2"
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
```

## License

MIT - Tomas Aparicio
