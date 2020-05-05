package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 为了示范反向代理而写的一个简单的下游服务器
func main() {
	rs1 := &RealServer{Addr: "127.0.0.1:7001"}
	rs1.Run()
	rs2 := &RealServer{Addr: "127.0.0.1:7002"}
	rs2.Run()

	//监听关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type RealServer struct {
	Addr string
}

func (r *RealServer) Run() {
	log.Println("Starting http Server at " + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HelloHandler)
	mux.HandleFunc("/base", r.HelloHandler)
	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Host)
	newPath := fmt.Sprintf("http://%s%s\n", r.Addr, req.URL.Path)
	realIP := fmt.Sprintf("RemoteAddr=%s\n", req.RemoteAddr)

	io.WriteString(w, newPath)
	io.WriteString(w, realIP)
}

func (r *RealServer) BaseHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is base")
}
