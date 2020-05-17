package main

import (
	"context"
	"fmt"
	"github/pibigstar/go-gateway/demo/proxy/reverse/tcp/server"
	"net"
)

type tcpHandler struct {
}

func (t *tcpHandler) ServeTCP(ctx context.Context, conn net.Conn) {
	conn.Write([]byte("Hello TCP World! \n"))
}

var addr = ":6001"

// tcp测试server端
// telnet 127.0.0.1 6001
func main() {
	fmt.Println("starting tcp server in ", addr)
	tcpServer := server.TcpServer{
		Addr:    addr,
		Handler: &tcpHandler{},
	}
	if err := tcpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
