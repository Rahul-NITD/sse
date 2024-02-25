package sse

import "net/http"

type ChatHandler struct {
	http.Handler
}

func NewChatHandler() *ChatHandler {
	r := http.NewServeMux()

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {})

	hdlr := &ChatHandler{}
	hdlr.Handler = r
	return hdlr
}
