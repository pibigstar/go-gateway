package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"

	"github/pibigstar/go-gateway/demo/proxy/reverse/https/server/ssl"
)

const addr = "127.0.0.1:4000"

func main() {
	log.Println("Starting https server at " + addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	// 将http1.1 升级为 http2
	if err := http2.ConfigureServer(server, &http2.Server{}); err != nil {
		panic(err)
	}
	// 将http 升级为 https
	if err := server.ListenAndServeTLS(ssl.Path("server.crt"), ssl.Path("server.key")); err != nil {
		panic(err)
	}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, fmt.Sprintf("https://%s%s", addr, req.URL.Path))
}
