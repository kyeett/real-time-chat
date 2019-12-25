package chatserver

import (
	"fmt"
	"io"

	chat "github.com/kyeett/real-time-chat/proto"
	"google.golang.org/grpc"
)

var _ chat.ChatServer = &Server{}

type Server struct {
	redis      *RedisService
	grpcServer *grpc.Server
}

func (s *Server) StartChat(stream chat.Chat_StartChatServer) error {
	fmt.Println("Connection initiated", s.redis.client.ClientID().Val())
	sub := s.redis.client.Subscribe("chatroom 1")

	go func() {
		for i := 0; i < 10; i++ {
			recv := <-sub.Channel()
			fmt.Printf("Received %s\n", recv)

			stream.Send(&chat.Message{
				Message: recv.String() + " + server stuff",
			})
		}
	}()

	for {
		in, err := stream.Recv()
		fmt.Println("Receive on grpc stream")
		if err == io.EOF {
			return fmt.Errorf("closed")
		}
		if err != nil {
			return fmt.Errorf("Failed to receive a note : %v", err)
		}
		fmt.Println("Send message on redis")

		s.redis.SendMessage(in.GetMessage())

		err = stream.Send(&chat.Message{
			Message: in.GetMessage() + " + server stuff",
		})
		fmt.Println("CHAT3")
	}

	return nil
}

func NewServer(ss *grpc.Server) *Server {

	s := Server{
		redis: NewRedisService(),
		// grpc:  NewGRPCService(),
	}

	return &s
}

func (s *Server) Stop() {
	s.redis.Stop()
	// s.grpcServer.Stop()
}
