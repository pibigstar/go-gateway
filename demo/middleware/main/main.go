package main

import (
	"fmt"
	"github/pibigstar/go-gateway/demo/balance"
	"github/pibigstar/go-gateway/demo/middleware"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var addr = "127.0.0.1:7000"

// 中间件测试
func main() {
	//初始化方法数组路由器
	sliceRouter := middleware.NewRouter()
	sliceRouter.Group("/base").Use(
		middleware.TraceLogSliceMW(), // trace 打印
		middleware.RateLimiter(),     // 限流
		middleware.CircuitMW(),       // 熔断

		// 实际业务处理函数
		func(c *middleware.RouterContext) {
			c.Rw.Write([]byte("test middleware"))
		})

	//请求到反向代理
	sliceRouter.Group("/").Use(middleware.TraceLogSliceMW(), func(c *middleware.RouterContext) {
		fmt.Println("reverseProxy")
		reverseProxy(c).ServeHTTP(c.Rw, c.Req)
	})
	// 设置熔断器
	middleware.SetHystrixConf(true)

	// 创建路由控制器
	routerHandler := middleware.NewRouterHandler(nil, sliceRouter)
	fmt.Println("tcp is staring...")
	log.Fatal(http.ListenAndServe(addr, routerHandler))
}

// 反向代理
var reverseProxy = func(c *middleware.RouterContext) http.Handler {
	wp := balance.LoadBalanceFactory(balance.WeightPollingType)
	wp.Add("http://127.0.0.1:7001", "10")
	wp.Add("http://127.0.0.1:7002", "20")

	return NewReverseProxyWithBalance(c, wp)
}

// 带负载均衡得反向代理
func NewReverseProxyWithBalance(c *middleware.RouterContext, balance balance.Balance) *httputil.ReverseProxy {
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
