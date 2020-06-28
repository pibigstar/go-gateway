package response

import "github.com/gogf/gf/net/ghttp"

type Response struct {
	Code  int         `json:"code"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func Error(r *ghttp.Request, err error) {
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
