package proxy

import (
	"context"
	"github/pibigstar/go-gateway/demo/balance"
	"io"
	"log"
	"net"
	"time"
)

// 返回一个带负载均衡的TCP反向代理
func NewTcpBalanceReverseProxy(ctx context.Context, lb balance.Balance) *TcpReverseProxy {
	return func() *TcpReverseProxy {
		nextAddr, err := lb.Get()
		if err != nil {
			panic("tcp proxy get next addr fail")
		}
		return &TcpReverseProxy{
			ctx:             ctx,
			Addr:            nextAddr,
			KeepAlivePeriod: time.Second,
			DialTimeout:     time.Second,
		}
	}()
}

//TCP反向代理
type TcpReverseProxy struct {
	ctx                  context.Context //单次请求单独设置
	Addr                 string
	KeepAlivePeriod      time.Duration //设置
	DialTimeout          time.Duration //设置超时时间
	DialContext          func(ctx context.Context, network, address string) (net.Conn, error)
	OnDialError          func(src net.Conn, dstDialErr error)
	ProxyProtocolVersion int
}

//传入上游 conn，在这里完成下游连接与数据交换
func (tp *TcpReverseProxy) ServeTCP(ctx context.Context, conn net.Conn) {
	//设置连接超时
	var cancel context.CancelFunc
	if tp.DialTimeout >= 0 {
		ctx, cancel = context.WithTimeout(ctx, tp.dialTimeout())
	}
	dst, err := tp.dialContext()(ctx, "tcp", tp.Addr)
	if cancel != nil {
		cancel()
	}
	if err != nil {
		tp.onDialError()(conn, err)
		return
	}

	defer func() { go dst.Close() }() //记得退出下游连接

	//设置dst的 keepAlive 参数,在数据请求之前
	if ka := tp.keepAlivePeriod(); ka > 0 {
		if c, ok := dst.(*net.TCPConn); ok {
			c.SetKeepAlive(true)
			c.SetKeepAlivePeriod(ka)
		}
	}
	errc := make(chan error, 1)
	go tp.proxyCopy(errc, conn, dst)
	go tp.proxyCopy(errc, dst, conn)
	<-errc
}

func (tp *TcpReverseProxy) onDialError() func(conn net.Conn, dstDialErr error) {
	if tp.OnDialError != nil {
		return tp.OnDialError
	}
	return func(src net.Conn, dstDialErr error) {
		log.Printf("tcpproxy: for incoming conn %v, error dialing %q: %v", src.RemoteAddr().String(), tp.Addr, dstDialErr)
		src.Close()
	}
}

func (tp *TcpReverseProxy) proxyCopy(errc chan<- error, dst, conn net.Conn) {
	_, err := io.Copy(dst, conn)
	errc <- err
}

func (tp *TcpReverseProxy) dialTimeout() time.Duration {
	if tp.DialTimeout > 0 {
		return tp.DialTimeout
	}
	return 10 * time.Second
}

func (tp *TcpReverseProxy) dialContext() func(ctx context.Context, network, address string) (net.Conn, error) {
	if tp.DialContext != nil {
		return tp.DialContext
	}
	return (&net.Dialer{
		Timeout:   tp.DialTimeout,     //连接超时
		KeepAlive: tp.KeepAlivePeriod, //设置连接的检测时长
	}).DialContext
}

func (tp *TcpReverseProxy) keepAlivePeriod() time.Duration {
	if tp.KeepAlivePeriod != 0 {
		return tp.KeepAlivePeriod
	}
	return time.Minute
}
