package sse_chat

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/r3labs/sse/v2"
)

func StartClient() func() {

	// client := sse.NewClient("http://localhost:8000" + "/room")

	// cxt, cancel := context.WithCancel(context.Background())

	fmt.Print("Enter stream uuid : ")
	var id string
	fmt.Scan(&id)

	if uuid.Validate(id) != nil {
		println("invalid id")
		return func() {}
	}

	client := sse.NewClient("http://localhost:8000/room")
	cxt, cancel := context.WithCancel(context.Background())

	go client.SubscribeWithContext(cxt, id, func(msg *sse.Event) {
		println("Msg:", string(msg.Data))
	})

	println("Client Started")

	return cancel
}
