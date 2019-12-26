package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	chatserver "github.com/kyeett/real-time-chat"
)

var connectAddr = flag.String("connect", ":8080", "<address>:<port> to listen on")

func main() {
	flag.Parse()

	c, err := chatserver.NewDefaultClient(*connectAddr)
	if err != nil {
		log.Fatalf("failed to create client: %v\n", err)
	}

	if err := c.Connect(); err != nil {
		log.Fatalf("failed to connect client: %v\n", err)
	}
	defer c.Stop()
	fmt.Println("Connection successful")

	for i := 0; i < 5; i++ {
		if err := c.Send("my msg" + strconv.Itoa(i)); err != nil {
			log.Fatalf("failed to connect client: %v\n", err)
		}
	}

	fmt.Println("Send successful successful")

	time.Sleep(10000 * time.Millisecond)

	// m, err := c.ReceiveMessage()
	// if err != nil {
	// 	log.Fatalf("failed to receive message: %v\n", err)
	// }
	// fmt.Printf("received: %q\n", *m)

}
