package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

// 一个简单的反向代理服务器

// 需要代理的地址
var proxyAddr = "http://127.0.0.1:7001"

func main() {
	http.HandleFunc("/", handler)
	log.Println("Start serving on port 7000")
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析代理地址，并更改请求体的协议和主机
	// 本来请求的是 127.0.0.1:7000 地址，我们将其改成 127.0.0.1:7001
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		return
	}
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host

	// 2. 请求下游服务器（server）
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	// 3. 把下游请求内容返回给上游
	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	bufio.NewReader(resp.Body).WriteTo(w)
}
