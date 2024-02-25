package sse_test

import (
	sse "blog-sse"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestServer(t *testing.T) {

	svr := httptest.NewServer(sse.NewChatHandler())

	res, err := http.Get(svr.URL + "/healthz")
	assert.NoError(t, err, "Could not make request to /healthz")
	assert.Equal[int](t, http.StatusOK, res.StatusCode)

	defer svr.Close()

}
