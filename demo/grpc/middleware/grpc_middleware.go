package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"

	"github/pibigstar/go-gateway/demo/balance"
	"github/pibigstar/go-gateway/demo/grpc/middleware/interceptor"
	"github/pibigstar/go-gateway/demo/public"
)

const (
	// 代理grpc服务器地址
	proxyServer = "localhost:5001"
	// 实际请求的下游grpc服务器地址
	grpcServer = "localhost:5000"
)

func main() {
	lis, err := net.Listen("tcp", proxyServer)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	wp := balance.LoadBalanceFactory(balance.WeightPollingType)
	wp.Add(grpcServer, "10")
	grpcProxyHandler := NewGrpcLoadBalanceHandler(wp)

	counter, _ := public.NewFlowCountService("local_app", time.Second)
	s := grpc.NewServer(
		//流式方法拦截
		grpc.ChainStreamInterceptor(interceptor.GrpcAuthStreamInterceptor, interceptor.GrpcFlowCountStreamInterceptor(counter)),
		//自定义codec
		grpc.CustomCodec(proxy.Codec()),
		//自定义代理全局回调
		grpc.UnknownServiceHandler(grpcProxyHandler))

	fmt.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
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
