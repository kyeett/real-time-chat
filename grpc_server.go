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
	defer sub.Close()

	go func() {
		fmt.Println("Starting receive channel")
		ch := sub.Channel()
		for recv := range ch {
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
			fmt.Println("EOF")
			return fmt.Errorf("closed")
		}
		if err != nil {
			err = fmt.Errorf("Failed to receive a note : %v", err)
			fmt.Println(err)
			return err
		}
		fmt.Println("Send message on redis")

		fmt.Println("Sending message", in.Message)
		s.redis.SendMessage(in.Message)

		// err = stream.Send(&chat.Message{
		// 	Message: in.GetMessage() + " + server stuff",
		// })
	}
}

func NewServer(ss *grpc.Server, redisURL string) (*Server, error) {

	redisService, err := NewRedisService(redisURL)
	if err != nil {
		return nil, err
	}

	s := Server{
		redis: redisService,
		// grpc:  NewGRPCService(),
	}

	return &s, nil
}

func (s *Server) Stop() {
	s.redis.Stop()
	// s.grpcServer.Stop()
}
