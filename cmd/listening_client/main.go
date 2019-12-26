package main

import (
	"fmt"
	"log"

	chatserver "github.com/kyeett/real-time-chat"
)

func main() {
	port := ":8080"

	c, err := chatserver.NewDefaultClient(port)
	if err != nil {
		log.Fatalf("failed to create client: %v\n", err)
	}

	if err := c.Connect(); err != nil {
		log.Fatalf("failed to connect client: %v\n", err)
	}

	for {
		m, err := c.ReceiveMessage()
		if err != nil {
			log.Fatalf("failed to receive message: %v\n", err)
		}
		fmt.Printf("received: %q\n", *m)
	}
}
