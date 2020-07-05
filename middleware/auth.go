package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/response"
	"github/pibigstar/go-gateway/utils/token"
)

// 用户权限校验
func Auth() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {
		if !r.IsFileRequest() {
			_, err := token.GetUserInfoFromSession(r)
			if err != nil {
				response.Error(r, err)
			}
		}
		r.Middleware.Next()
	}
}
