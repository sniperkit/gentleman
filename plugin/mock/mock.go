package mock

import (
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugin"
	"gopkg.in/h2non/gock.v1"
)

// Plugin exports the mock plugin
var Plugin = plugin.NewPhasePlugin("before dial", func(ctx *context.Context, h context.Handler) {
	gock.InterceptClient(ctx.Client)
	h.Next(ctx)
})

// New creates a new gock mock.
// It's a shorthand to gock.New().
func New(uri string) *gock.Request {
	return gock.New(uri)
}

// Disable disables the registered mocks.
// It's a shorthand to gock.Disable().
func Disable() {
	gock.Disable()
}
