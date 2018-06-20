package mock

import (
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gentleman.v2"
)

func TestMock(t *testing.T) {
	defer Disable()

	New("http://foo.com").Reply(204).SetHeader("Server", "gock")

	req := gentleman.NewRequest()
	req.Use(Plugin)
	req.SetHeader("foo", "bar")
	req.URL("http://foo.com")

	res, err := req.Send()
	st.Expect(t, err, nil)
	st.Expect(t, res.Ok, true)
	st.Expect(t, res.StatusCode, 204)
	st.Expect(t, res.Header.Get("Server"), "gock")
}

func TestMockMatchHeader(t *testing.T) {
	defer Disable()

	New("http://foo.com").MatchHeader("foo", "bar").Reply(204).SetHeader("Server", "gock")

	req := gentleman.NewRequest()
	req.Use(Plugin)
	req.SetHeader("foo", "bar")
	req.URL("http://foo.com")

	res, err := req.Send()
	st.Expect(t, err, nil)
	st.Expect(t, res.Ok, true)
	st.Expect(t, res.StatusCode, 204)
	st.Expect(t, res.Header.Get("Server"), "gock")
}

func TestMockError(t *testing.T) {
	defer Disable()

	New("http://bar.com").Reply(204).SetHeader("Server", "gock")

	req := gentleman.NewRequest()
	req.Use(Plugin)
	req.SetHeader("foo", "bar")
	req.URL("http://foo.com")

	res, err := req.Send()
	st.Reject(t, err, nil)
	st.Expect(t, res.Ok, false)
	st.Expect(t, res.StatusCode, 0)
}
