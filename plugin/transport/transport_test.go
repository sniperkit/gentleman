package transport

import (
	"net/http"
	"testing"

	"github.com/nbio/st"

	"github.com/sniperkit/gentleman/pkg/context"
)

func TestSetTransport(t *testing.T) {
	ctx := context.New()
	fn := newHandler()
	transport := &http.Transport{}
	Set(transport).Exec("request", ctx, fn.fn)
	st.Expect(t, fn.called, true)
	newTransport := ctx.Client.Transport.(*http.Transport)
	st.Expect(t, newTransport, transport)
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