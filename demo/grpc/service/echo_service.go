package service

import (
	"context"
	"fmt"
	"github/pibigstar/go-gateway/demo/grpc/pb/echo"
	"io"
)

const (
	streamingCount = 10
)

type EchoService struct {
	*echo.UnimplementedEchoServer
}

func (s *EchoService) UnaryEcho(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	fmt.Printf("request received: %v, sending echo\n", in)
	return &echo.EchoResponse{Message: in.Message}, nil
}

func (s *EchoService) ServerStreamingEcho(in *echo.EchoRequest, stream echo.Echo_ServerStreamingEchoServer) error {
	fmt.Printf("--- ServerStreamingEcho ---\n")
	fmt.Printf("request received: %v\n", in)
	// Read requests and send responses.
	for i := 0; i < streamingCount; i++ {
		fmt.Printf("echo message %v\n", in.Message)
		err := stream.Send(&echo.EchoResponse{Message: in.Message})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *EchoService) ClientStreamingEcho(stream echo.Echo_ClientStreamingEchoServer) error {
	fmt.Printf("--- ClientStreamingEcho ---\n")
	// Read requests and send responses.
	var message string
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("echo last received message\n")
			return stream.SendAndClose(&echo.EchoResponse{Message: message})
		}
		message = in.Message
		fmt.Printf("request received: %v, building echo\n", in)
		if err != nil {
			return err
		}
	}
}

func (s *EchoService) BidirectionalStreamingEcho(stream echo.Echo_BidirectionalStreamingEchoServer) error {
	fmt.Printf("--- BidirectionalStreamingEcho ---\n")
	// Read requests and send responses.
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("request received %v, sending echo\n", in)
		if err := stream.Send(&echo.EchoResponse{Message: in.Message}); err != nil {
			return err
		}
	}
}
