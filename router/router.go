package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/api/user"
	"github/pibigstar/go-gateway/middleware"
)

func init() {
	s := g.Server()
	s.Use(middleware.Trace())
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/login", user.Login)
	})
}
