package main

import (
	"context"
	"fmt"
	"github.com/e421083458/grpc-proxy/proxy"
	"github/pibigstar/go-gateway/demo/balance"
	"google.golang.org/grpc"
	"log"
	"net"
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
	wp := balance.LoadBalanceFactory(balance.WeightPollingType)
	wp.Add(grpcServer, "10")

	grpcProxyHandler := NewGrpcLoadBalanceHandler(wp)
	server := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(grpcProxyHandler))

	fmt.Printf("grpc proxy server listening at %v\n", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewGrpcLoadBalanceHandler(balance balance.Balance) grpc.StreamHandler {
	return func() grpc.StreamHandler {
		nextAddr, err := balance.Get()
		if err != nil {
			log.Fatal("get next addr fail")
		}
		director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
			c, err := grpc.DialContext(ctx, nextAddr, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
			return ctx, c, err
		}
		return proxy.TransparentHandler(director)
	}()
}
