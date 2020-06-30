package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	_ "github/pibigstar/go-gateway/boot"
	_ "github/pibigstar/go-gateway/router"
)

func main() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
	s.Run()
}
