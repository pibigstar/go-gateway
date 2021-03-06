package main

import (
	"context"
	"fmt"
	"github/pibigstar/go-gateway/demo/grpc/pb/echo"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	timestampFormat = time.StampNano // "Jan _2 15:04:05.000"
	streamingCount  = 10
	message         = "this is examples/metadata"
	addr            = "localhost:5000" // grpc实际服务地址
	proxyAddr       = "localhost:5001" // grpc代理服务器地址
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := echo.NewEchoClient(conn)

	//调用一元方法
	unaryCallWithMetadata(client, message)

	////服务端流式
	//serverStreamingWithMetadata(client, message)
	//
	////客户端流式
	//clientStreamWithMetadata(client, message)
	//
	////双向流式
	//bidirectionalWithMetadata(client, message)

	time.Sleep(3 * time.Second)
}

func unaryCallWithMetadata(client echo.EchoClient, message string) {
	fmt.Printf("--- unary ---\n")

	// Create metadata and context.
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer some-secret-token")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		r, err := client.UnaryEcho(ctx, &echo.EchoRequest{Message: message})
		if err != nil {
			log.Fatalf("failed to call UnaryEcho: %v", err)
		}
		fmt.Printf("response:%v\n", r.Message)
	}()
	cancel()
	return
}

func serverStreamingWithMetadata(c echo.EchoClient, message string) {
	fmt.Printf("--- server streaming ---\n")

	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer some-secret-token")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := c.ServerStreamingEcho(ctx, &echo.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("failed to call ServerStreamingEcho: %v", err)
	}

	// Read all the responses.
	var rpcStatus error
	fmt.Printf("response:\n")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", r.Message)
	}
	if rpcStatus != io.EOF {
		log.Fatalf("failed to finish server streaming: %v", rpcStatus)
	}
}

func clientStreamWithMetadata(c echo.EchoClient, message string) {
	fmt.Printf("--- client streaming ---\n")
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer some-secret-token")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := c.ClientStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("failed to call ClientStreamingEcho: %v\n", err)
	}

	// Send all requests to the server.
	for i := 0; i < streamingCount; i++ {
		if err := stream.Send(&echo.EchoRequest{Message: message}); err != nil {
			log.Fatalf("failed to send streaming: %v\n", err)
		}
	}

	// Read the response.
	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to CloseAndRecv: %v\n", err)
	}
	fmt.Printf("response:%v\n", r.Message)
}

func bidirectionalWithMetadata(c echo.EchoClient, message string) {
	fmt.Printf("--- bidirectional ---\n")
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md.Append("authorization", "Bearer some-secret-token")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := c.BidirectionalStreamingEcho(ctx)
	if err != nil {
		log.Fatalf("failed to call BidirectionalStreamingEcho: %v\n", err)
	}

	go func() {
		// Send all requests to the server.
		for i := 0; i < streamingCount; i++ {
			if err := stream.Send(&echo.EchoRequest{Message: message}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// Read all the responses.
	var rpcStatus error
	fmt.Printf("response:\n")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", r.Message)
	}
	if rpcStatus != io.EOF {
		log.Fatalf("failed to finish server streaming: %v", rpcStatus)
	}
}
