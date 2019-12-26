package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/kyeett/chatui"
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

	ui, err := chatui.New()
	if err != nil {
		log.Fatal(err)
	}

	ui.SetInputCallback(func(p chatui.Post) {
		ui.AppendNewPost(p)
		if err := c.Send(p.Username + " " + p.Message); err != nil {

			// log.Fatalf("failed to connect client: %v\n", err)
		}
	})

	// Simulate receiving messages
	go func() {
		for {
			m, err := c.ReceiveMessage()
			if err != nil {
				log.Fatalf("failed to receive message: %v\n", err)
			}

			ui.AppendNewPost(chatui.Post{
				Username: "someone",
				Message:  *m,
				Time:     time.Now().Format("15:04"),
			})

		}
	}()

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}

	// for i := 0; i < 5; i++ {
	// 	if err := c.Send("my msg" + strconv.Itoa(i)); err != nil {
	// 		log.Fatalf("failed to connect client: %v\n", err)
	// 	}
	// }

	fmt.Println("Send successful successful")

	// m, err := c.ReceiveMessage()
	// if err != nil {
	// 	log.Fatalf("failed to receive message: %v\n", err)
	// }
	// fmt.Printf("received: %q\n", *m)

}
