package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"runtime/debug"
)

// Recovery捕获所有panic
func Recovery() func(r *ghttp.Request)   {
	return func(r *ghttp.Request) {
		defer func() {
			if err := recover(); err != nil {
				glog.Println(string(debug.Stack()))
			}
		}()
		r.Middleware.Next()
	}
}
