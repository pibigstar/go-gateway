package main

import (
	"github/pibigstar/go-gateway/demo/balance"
	"github/pibigstar/go-gateway/demo/proxy/reverse/https/proxy"
	"github/pibigstar/go-gateway/demo/proxy/reverse/https/server/ssl"
	"log"
	"net/http"
)

var addr = "127.0.0.1:4001"

func main() {
	rb := balance.LoadBalanceFactory(balance.WeightPollingType)
	rb.Add("https://127.0.0.1:4000", "50")

	log.Println("Starting https server at " + addr)

	// https 反向代理
	httpsProxy := proxy.NewMultipleHostsReverseProxy(rb)
	log.Fatal(http.ListenAndServeTLS(addr, ssl.Path("server.crt"), ssl.Path("server.key"), httpsProxy))
}
