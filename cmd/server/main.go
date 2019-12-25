package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	chatserver "github.com/kyeett/real-time-chat"
	"google.golang.org/grpc"

	chat "github.com/kyeett/real-time-chat/proto"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	port := ":8901"

	ss := grpc.NewServer()
	s := chatserver.NewServer(ss)
	chat.RegisterChatServer(ss, s)

	go func() {
		<-c
		fmt.Printf("Received signal, shutting down application\n")
		ss.Stop()
		s.Stop()
	}()

	fmt.Printf("Listening on port %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	if err := ss.Serve(lis); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Application shut down")
}
