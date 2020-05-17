package server

import (
	"context"
	"fmt"
	"net"
	"runtime"
	"sync"
)

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	return tc, nil
}

type onceCloseListener struct {
	net.Listener
	once     sync.Once
	closeErr error
}

func (oc *onceCloseListener) Close() error {
	oc.once.Do(oc.close)
	return oc.closeErr
}

func (oc *onceCloseListener) close() { oc.closeErr = oc.Listener.Close() }

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "tcp_proxy context value " + k.name }

type conn struct {
	server     *TcpServer
	cancelCtx  context.CancelFunc
	conn       net.Conn
	remoteAddr string
}

func (c *conn) close() {
	c.conn.Close()
}

// 连接成功后的处理函数
func (c *conn) serve(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil && err != ErrAbortHandler {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			fmt.Printf("tcp: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
		c.close()
	}()

	c.remoteAddr = c.conn.RemoteAddr().String()
	ctx = context.WithValue(ctx, LocalAddrContextKey, c.conn.LocalAddr())
	if c.server.Handler == nil {
		panic("handler empty")
	}
	c.server.Handler.ServeTCP(ctx, c.conn)
}
