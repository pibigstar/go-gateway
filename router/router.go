package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/api/admin"
	"github/pibigstar/go-gateway/app/api/gateway"
	"github/pibigstar/go-gateway/middleware"
)

func init() {
	s := g.Server()

	s.Use(middleware.Recovery())
	s.Use(middleware.Trace())
	s.Use(middleware.IPLimit())

	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.POST("/login", admin.Login)
		group.GET("/info", admin.Info)
		group.GET("/logout", admin.Logout)
		group.POST("/changePwd", admin.ChangePwd)
	})

	s.Group("/gateway", func(group *ghttp.RouterGroup) {
		// 中间件，登陆校验
		group.Middleware(middleware.Auth())

		group.GET("/list", gateway.List)
		group.GET("/detail", gateway.Detail)
		group.GET("/stat", gateway.Stat)
	})
}
