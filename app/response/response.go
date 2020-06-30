package response

import (
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/utils/errx"
)

type Response struct {
	Code  int         `json:"code"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func Error(r *ghttp.Request, err error) {
	if e, ok := err.(errx.ErrorX); ok {
		_ = r.Response.WriteJsonExit(Response{
			Code:  e.Code(),
			Error: e.Error(),
		})
		return
	}
	_ = r.Response.WriteJsonExit(Response{
		Code:  500,
		Error: err.Error(),
	})
}

func Success(r *ghttp.Request, data interface{}) {
	_ = r.Response.WriteJsonExit(Response{
		Code: 200,
		Data: data,
	})
}
