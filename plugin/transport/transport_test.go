package transport

import (
	"net/http"
	"testing"

	"github.com/nbio/st"

	c "github.com/sniperkit/gentleman/pkg/context"
)

func TestSetTransport(t *testing.T) {
	ctx := c.New()
	fn := newHandler()
	transport := &http.Transport{}
	Set(transport).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	newTransport := ctx.Client.Transport.(*http.Transport)
	st.Expect(t, newTransport, transport)
}

type handler struct {
	fn     c.Handler
	called bool
}

func newHandler() *handler {
	h := &handler{}
	h.fn = c.NewHandler(func(c *c.Context) {
		h.called = true
	})
	return h
}
