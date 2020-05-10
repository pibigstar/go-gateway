package main

import (
	"github/pibigstar/go-gateway/demo/balance"
	"github/pibigstar/go-gateway/demo/proxy/reverse/http/proxy"
	"log"
	"net/http"
)

var (
	addr = "127.0.0.1:7004"
)

func main() {
	// 访问 7004, 反向代理到7003
	rb := balance.LoadBalanceFactory(balance.WeightPollingType)
	rb.Add("http://127.0.0.1:7003", "50")

	reverseProxy := proxy.NewReverseProxyWithBalance(rb)
	log.Println("Starting http server at " + addr)
	log.Fatal(http.ListenAndServe(addr, reverseProxy))
}
