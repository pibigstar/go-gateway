package main

import (
	"context"
	"fmt"
	"github/pibigstar/go-gateway/demo/grpc/pb/echo"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	serverAddr     = ":8081"
	grpcServerAddr = "localhost:5000"
)

func main() {

	fmt.Println("server listening at", serverAddr)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := echo.RegisterEchoHandlerFromEndpoint(ctx, mux, grpcServerAddr, opts)
	if err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		panic(err)
	}
}
