package logger

import (
	"io"
	"net/http"

	"github.com/izumin5210/httplogger"

	c "github.com/sniperkit/gentleman/pkg/context"
	p "github.com/sniperkit/gentleman/pkg/plugin"
)

// New creates logger plugin instance
func New(out io.Writer) p.Plugin {
	return new(func(parent http.RoundTripper) http.RoundTripper {
		return httplogger.NewRoundTripper(out, parent)
	})
}

// FromLogger creates logger plugin instance with a specified logger implementation
func FromLogger(writer httplogger.SimpleLogWriter) p.Plugin {
	return new(func(parent http.RoundTripper) http.RoundTripper {
		return httplogger.FromSimpleLogger(writer, parent)
	})
}

func new(transportFn func(parent http.RoundTripper) http.RoundTripper) p.Plugin {
	return p.NewRequestPlugin(func(c *c.Context, h c.Handler) {
		c.Client.Transport = transportFn(c.Client.Transport)
		h.Next(c)
	})
}
