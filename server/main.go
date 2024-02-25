package main

import (
	sse_chat "blog-sse"
	"log"
	"net/http"
)

func main() {
	log.Fatal(
		http.ListenAndServe("localhost:8000", sse_chat.NewChatHandler()),
	)
}
