package main

import (
	"github.com/gogf/gf/frame/g"
	_ "github/pibigstar/go-gateway/boot"
	_ "github/pibigstar/go-gateway/router"
)

func main() {
	g.Server().Run()
}
