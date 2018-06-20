package compression

import (
	"net/http"
	"testing"

	"github.com/nbio/st"

	"github.com/sniperkit/gentleman/pkg/context"
)

func TestDisableCompression(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	Disable().Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	transport := ctx.Client.Transport.(*http.Transport)
	st.Expect(t, transport.DisableCompression, true)
}

type handler struct {
	fn     context.Handler
	called bool
}

func newHandler() *handler {
	h := &handler{}
	h.fn = context.NewHandler(func(c *context.Context) {
		h.called = true
	})
	return h
}
