package transport

import (
	"net/http"

	c "github.com/sniperkit/gentleman/pkg/context"
	p "github.com/sniperkit/gentleman/pkg/plugin"
)

// Set sets a new HTTP transport for the outgoing request
func Set(transport http.RoundTripper) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		// Override the http.Client transport
		ctx.Client.Transport = transport
		h.Next(ctx)
	})
}
