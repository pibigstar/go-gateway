package balace

import (
	"github/pibigstar/go-gateway/demo/balance"
	"github/pibigstar/go-gateway/demo/proxy/reverse/http/proxy"
	"net/http"
)

// 带负载均衡的反向代理
func main() {
	// 加权负载均衡
	wp := balance.LoadBalanceFactory(balance.WeightPollingType)
	wp.Add("http://127.0.0.1:7001", "10")
	wp.Add("http://127.0.0.1:7002", "20")

	// 拥有了加权负载的控制器
	reverseProxy := proxy.NewReverseProxyWithBalance(wp)

	if err := http.ListenAndServe(":7000", reverseProxy); err != nil {
		panic(err)
	}
}
