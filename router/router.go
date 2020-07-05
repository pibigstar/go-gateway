package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/api/admin"
	"github/pibigstar/go-gateway/app/api/service"
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

	s.Group("/service", func(group *ghttp.RouterGroup) {
		//group.Hook("/*", ghttp.HOOK_BEFORE_SERVE, middleware.Auth())
		group.Middleware(middleware.Auth())
		group.GET("/list", service.List)
	})
}
