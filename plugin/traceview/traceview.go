package traceview

import (
	"github.com/tracelytics/go-traceview/v1/tv"
	"golang.org/x/net/context"

	c "github.com/sniperkit/gentleman/pkg/context"
	p "github.com/sniperkit/gentleman/pkg/plugin"
)

// New creates gentleman plugin for TraceView http client.
func New(ctx context.Context) p.Plugin {
	p := plugin.New()
	p.SetHandler("request", func(gctx *c.Context, h c.Handler) {
		l := tv.BeginHTTPClientLayer(ctx, gctx.Request)
		defer func() {
			l.AddHTTPResponse(gctx.Response, gctx.Error)
			l.End()
		}()
		h.Next(gctx)
	})
	return p
}
