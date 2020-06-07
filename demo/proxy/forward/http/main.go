package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

// 正向代理服务器
// 可以将web请求在发送到服务器之前，我们对其请求做一些改写
func main() {
	fmt.Println("Forward Serve is running...")
	http.Handle("/", &ForwardProxy{})
	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}

type ForwardProxy struct{}

func (*ForwardProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)

	// 1. 复制一个新的req，并设置一些 Header头
	newReq := &http.Request{}
	*newReq = *req // 浅拷贝
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := newReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		newReq.Header.Set("X-Forwarded-For", clientIP)
		newReq.Header.Set("X-My-info", "pibigstar")
	}

	// 2. 请求下游数据
	res, err := http.DefaultTransport.RoundTrip(newReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	// 3. 把下游请求内容返回给上游
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
	res.Body.Close()
}
