package main

import (
	"context"
	"fmt"
	"github/pibigstar/go-gateway/demo/balance"
	"github/pibigstar/go-gateway/demo/proxy/reverse/tcp/proxy"
	"github/pibigstar/go-gateway/demo/proxy/reverse/tcp/server"
)

var addr = ":6000"

// 测试tcp反向代理到redis服务器
// telnet 127.0.0.1 6000
func main() {
	fmt.Println("starting tcp proxy tcp in ", addr)

	wp := balance.LoadBalanceFactory(balance.WeightPollingType)
	wp.Add("127.0.0.1:6379", "10")
	tcpProxy := proxy.NewTcpBalanceReverseProxy(context.Background(), wp)

	tcpServer := server.TcpServer{
		Addr:    addr,
		Handler: tcpProxy,
	}
	if err := tcpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
