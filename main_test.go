package chatserver

import (
	"github.com/stretchr/testify/require"
	"fmt"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"

	chat "github.com/kyeett/real-time-chat/proto"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	port := ":8901"

	ss := grpc.NewServer()
	s, err := NewServer(ss, "redis://127.0.0.1:6379")
	assert.NoError(t, err)
	chat.RegisterChatServer(ss, s)

	// Defer clean up
	defer ss.Stop()
	defer s.Stop()

	lis, err := net.Listen("tcp", port)
	require.NoError(t,err)
	
	go ss.Serve(lis)

	c1, err := NewDefaultClient(port)
	assert.NoError(t, err)
	c1.Connect()

	// c2, err := NewDefaultClient(port)
	// assert.NoError(t, err)
	// c2.Start()

	c1.Send("kaiyu")
	c1.Send("kaiyu2")
	m, err := c1.stream.Recv()
	assert.NoError(t, err)
	fmt.Printf("Received on stream %q\n", m.GetMessage())
	// c1.Send("kaiyu")

	time.Sleep(300 * time.Millisecond)

	//  c2.Receive()

	assert.NoError(t, err)

	// for _, s := range []string{"abc", "xyz"} {
	// 	stream.Send(&chat.Message{
	// 		Message: s,
	// 	})

	// 	m, err := stream.Recv()
	// 	if err != nil {
	// 		return
	// 	}
	// 	fmt.Println(m)
	// }
}

/*
stream, err := client.RouteChat(context.Background())
waitc := make(chan struct{})
go func() {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			// read done.
			close(waitc)
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive a note : %v", err)
		}
		log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
	}
}()
for _, note := range notes {
	if err := stream.Send(note); err != nil {
		log.Fatalf("Failed to send a note: %v", err)
	}
}
stream.CloseSend()
<-waitc
*/
