package chatserver

import (
	"context"

	chat "github.com/kyeett/real-time-chat/proto"
	"google.golang.org/grpc"
)

type client struct {
	grpcClient chat.ChatClient
	stream     chat.Chat_StartChatClient
}

func NewDefaultClient(port string) (*client, error) {
	opts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	conn, err := grpc.Dial(port, opts...)
	if err != nil {
		return nil, err
	}
	c := chat.NewChatClient(conn)
	return &client{grpcClient: c}, nil
}

func (c *client) Connect() error {
	stream, err := c.grpcClient.StartChat(context.Background())
	if err != nil {
		return err
	}
	c.stream = stream
	return nil
}

func (c *client) Send(msg string) error {
	return c.stream.Send(&chat.Message{
		Message: msg,
	})
}

func (c *client) ReceiveMessage() (*string, error) {
	m, err := c.stream.Recv()
	if err != nil {
		return nil, err
	}
	s := m.GetMessage()
	return &s, nil
}

func (c *client) Stop() {
	if c.stream == nil {
		return
	}

	c.stream.CloseSend()
}
