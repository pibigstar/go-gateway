package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrServerClosed     = errors.New("tcp: Server closed")
	ErrAbortHandler     = errors.New("tcp: abort TCPHandler")
	ServerContextKey    = &contextKey{"tcp-tcp"}
	LocalAddrContextKey = &contextKey{"local-addr"}
)

type TCPHandler interface {
	ServeTCP(ctx context.Context, conn net.Conn)
}

type TcpServer struct {
	Addr    string
	Handler TCPHandler
	err     error
	BaseCtx context.Context

	WriteTimeout     time.Duration
	ReadTimeout      time.Duration
	KeepAliveTimeout time.Duration

	DialTimeout time.Duration
	//下游参数
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

	mu         sync.Mutex
	inShutdown int32
	doneChan   chan struct{}
	onceCloser *onceCloseListener
}

func (ts *TcpServer) shuttingDown() bool {
	return atomic.LoadInt32(&ts.inShutdown) != 0
}

func (ts *TcpServer) ListenAndServe() error {
	if ts.shuttingDown() {
		return ErrServerClosed
	}
	addr := ts.Addr
	if addr == "" {
		return errors.New("need addr")
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return ts.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}

func (ts *TcpServer) Close() error {
	atomic.StoreInt32(&ts.inShutdown, 1)
	close(ts.doneChan)
	return ts.onceCloser.Close()
}

func (ts *TcpServer) Serve(l net.Listener) error {
	ts.onceCloser = &onceCloseListener{Listener: l}
	defer ts.onceCloser.Close()

	if ts.BaseCtx == nil {
		ts.BaseCtx = context.Background()
	}
	baseCtx := ts.BaseCtx
	ctx := context.WithValue(baseCtx, ServerContextKey, ts)
	for {
		conn, e := l.Accept()
		if e != nil {
			select {
			case <-ts.getDoneChan():
				return ErrServerClosed
			default:
			}
			return e
		}
		// 创建一个带各种参数的tcp连接
		c := ts.newConnWithTimeOut(conn)
		// 连接成功后读取信息，处理逻辑
		go c.serve(ctx)
	}
	return nil
}

func (ts *TcpServer) newConnWithTimeOut(rwc net.Conn) *conn {
	c := &conn{
		server: ts,
		conn:   rwc,
	}
	// 设置参数
	if d := c.server.ReadTimeout; d != 0 {
		if err := c.conn.SetReadDeadline(time.Now().Add(d)); err != nil {
			fmt.Println("set read deadline, err:", err.Error())
		}
	}
	if d := c.server.WriteTimeout; d != 0 {
		if err := c.conn.SetWriteDeadline(time.Now().Add(d)); err != nil {
			fmt.Println("set write deadline, err:", err.Error())
		}
	}
	if d := c.server.KeepAliveTimeout; d != 0 {
		if tcpConn, ok := c.conn.(*net.TCPConn); ok {
			if err := tcpConn.SetKeepAlive(true); err != nil {
				fmt.Println("set keep alive, err:", err.Error())
			}
			if err := tcpConn.SetKeepAlivePeriod(d); err != nil {
				fmt.Println("set keep alive period, err:", err.Error())
			}
		}
	}
	return c
}

func (ts *TcpServer) getDoneChan() <-chan struct{} {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	return ts.getDoneChanLocked()
}

func (ts *TcpServer) getDoneChanLocked() chan struct{} {
	if ts.doneChan == nil {
		ts.doneChan = make(chan struct{})
	}
	return ts.doneChan
}

func ListenAndServe(addr string, handler TCPHandler) error {
	server := &TcpServer{Addr: addr, Handler: handler, doneChan: make(chan struct{})}
	return server.ListenAndServe()
}
