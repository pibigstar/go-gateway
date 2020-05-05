package main

import (
	"fmt"
	"github/pibigstar/go-gateway/demo/proxy/balance"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// 带负载均衡的反向代理
func main() {
	// 加权负载均衡
	wp := balance.LoadBalanceFactory(balance.WeightPollingType)
	wp.Add("http://127.0.0.1:7001", "10")
	wp.Add("http://127.0.0.1:7002", "20")

	// 拥有了加权负载的控制器
	reverseProxy := NewReverseProxyWithBalance(wp)

	if err := http.ListenAndServe(":7000", reverseProxy); err != nil {
		panic(err)
	}
}

func NewReverseProxyWithBalance(balance balance.Balance) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		targetAddr, err := balance.Get()
		if err != nil {
			fmt.Printf("failed to get url, err: %v", err)
		}
		target, err := url.Parse(targetAddr)
		if err != nil {
			fmt.Printf("failed to parse url, err: %v", err)
		}
		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	return &httputil.ReverseProxy{Director: director}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
