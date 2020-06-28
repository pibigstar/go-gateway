package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/api/admin"
	"github/pibigstar/go-gateway/middleware"
)

func init() {
	s := g.Server()

	s.Use(middleware.Recovery())
	s.Use(middleware.Trace())
	s.Use(middleware.IPLimit())

	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.GET("/login", admin.Login)
	})
}
