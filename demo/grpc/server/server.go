package main

import (
	"fmt"
	"github/pibigstar/go-gateway/demo/grpc/pb/echo"
	"github/pibigstar/go-gateway/demo/grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	server := grpc.NewServer()
	echo.RegisterEchoServer(server, &service.EchoService{})
	server.Serve(lis)
}
