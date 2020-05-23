package main

import (
	"context"
	"fmt"
	"github.com/e421083458/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strings"
)

const (
	// 代理grpc服务器地址
	proxyServer = "localhost:5001"
	// 实际请求的下游grpc服务器地址
	grpcServer = "localhost:5000"
)

// grpc代理服务器
func main() {
	listener, err := net.Listen("tcp", proxyServer)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		// 限制访问内部接口
		if strings.Contains(fullMethodName, "internal") {
			return ctx, nil, status.Errorf(codes.Unimplemented,
				"Unknown method")
		}
		conn, err := grpc.DialContext(ctx, grpcServer, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
		return ctx, conn, err
	}

	server := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)))

	fmt.Printf("grpc proxy server listening at %v\n", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
