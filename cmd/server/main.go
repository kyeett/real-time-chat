package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	chatserver "github.com/kyeett/real-time-chat"
	"google.golang.org/grpc"

	chat "github.com/kyeett/real-time-chat/proto"
)

var listenAddr = flag.String("listen", ":8080", "<address>:<port> to listen on")

func main() {
	flag.Parse()

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("Missing REDIS_URL")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ss := grpc.NewServer()
	s, err := chatserver.NewServer(ss, redisURL)
	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
	chat.RegisterChatServer(ss, s)

	go func() {
		<-c
		fmt.Printf("Received signal, shutting down application\n")
		ss.Stop()
		s.Stop()
	}()

	fmt.Printf("Listening on %s\n", *listenAddr)
	lis, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	if err := ss.Serve(lis); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Application shut down")
}
