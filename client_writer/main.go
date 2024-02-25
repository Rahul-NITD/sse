package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func main() {
	var id string

	fmt.Print("Enter room id: ")
	fmt.Scanln(&id)
	if uuid.Validate(id) != nil {
		println("Invalid id")
		return
	}

	fmt.Println("Start writing messages")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		res, err := http.Post("http://localhost:8000/dump?id="+id, "text/plain", strings.NewReader(msg))
		if err != nil {
			panic(err)
		}
		if res.StatusCode != http.StatusOK {
			println("Could not send message to server, got", res.StatusCode)
		}
	}
}
