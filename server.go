package sse_chat

import (
	"io"
	"net/http"

	"github.com/r3labs/sse/v2"
)

const (
	HEALTHZ_PATH    = "/healthz"
	ROOM_PATH       = "/room"
	DUMP_PATH       = "/dump"
	ID_PARAM        = "id"
	DEFAULT_WELCOME = "Hello User"
)

func NewChatHandler() http.Handler {
	r := http.NewServeMux()

	sse_server := CreateSSEServer()

	r.HandleFunc(HEALTHZ_PATH, handleHealthZ)
	r.HandleFunc(ROOM_PATH, sse_server.ServeHTTP)
	r.HandleFunc(DUMP_PATH, handleDumpDecorator(sse_server))

	return r
}

func CreateSSEServer() *sse.Server {
	sse_server := sse.New()
	sse_server.OnSubscribe = OnSubscribeDecorator(sse_server)
	sse_server.AutoStream = true
	return sse_server
}

func handleHealthZ(w http.ResponseWriter, r *http.Request) {}

func OnSubscribeDecorator(sse_server *sse.Server) func(streamID string, sub *sse.Subscriber) {
	return func(streamID string, sub *sse.Subscriber) {
		sse_server.Publish(streamID, &sse.Event{
			Data: []byte(DEFAULT_WELCOME),
		})
	}
}

func handleDumpDecorator(sse_server *sse.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := io.ReadAll(r.Body)
		sse_server.Publish(r.URL.Query().Get(ID_PARAM), &sse.Event{
			Data: body,
		})
	}
}
