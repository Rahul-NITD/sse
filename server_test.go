package sse_chat_test

import (
	sse_chat "blog-sse"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/r3labs/sse/v2"
)

func TestServer(t *testing.T) {

	svr := httptest.NewServer(sse_chat.NewChatHandler())

	res, err := http.Get(svr.URL + "/healthz")
	assert.NoError(t, err, "Could not make request to /healthz")
	assert.Equal[int](t, http.StatusOK, res.StatusCode)

	defer svr.Close()

}

func TestSSEServer(t *testing.T) {

	svr := httptest.NewServer(sse_chat.NewChatHandler())

	client := sse.NewClient(svr.URL + "/room")

	cxt, cancel := context.WithCancel(context.Background())

	msgChan := make(chan []byte)
	go client.SubscribeWithContext(cxt, "550e8400-e29b-41d4-a716-446655440000", func(msg *sse.Event) {
		msgChan <- msg.Data
	})

	select {
	case <-time.After(100 * time.Millisecond):
		t.Error("Expected a message. Deadline expired.")
	case msg := <-msgChan:
		assert.Equal[[]byte](t, []byte("Hello User"), msg)
	}
	cancel()
	svr.Close()
}
