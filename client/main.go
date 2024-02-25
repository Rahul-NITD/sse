package main

import (
	sse_chat "blog-sse"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cancel := sse_chat.StartClient()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	cancel()
	println("Client closed")
}
