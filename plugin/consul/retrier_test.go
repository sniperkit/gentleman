package consul

import (
	"testing"

	"github.com/nbio/st"

	c "github.com/sniperkit/gentleman/pkg/context"
)

func TestRetrier(t *testing.T) {
	consul := NewClient(NewConfig("consul.server", "foo"))
	retrier := &Retrier{Consul: consul, Context: c.New(), Retry: DefaultRetrier}

	calls := 0
	retrier.Run(func() error {
		calls++
		return nil
	})

	st.Expect(t, calls, 1)
}

func TestNewRetrier(t *testing.T) {
	consul := NewClient(NewConfig("consul.server", "foo"))
	retrier := NewRetrier(consul, c.New())

	calls := 0
	retrier.Run(func() error {
		calls++
		return nil
	})

	st.Expect(t, calls, 1)
}
