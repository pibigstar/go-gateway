package admin

import (
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/model"
	"github/pibigstar/go-gateway/app/request"
	"github/pibigstar/go-gateway/app/response"
)

func Login(r *ghttp.Request) {
	var req *request.AdminLoginReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err)
	}

	resp, err := model.MAdminModel.AdminLogin(r.Context(), req)
	if err != nil {
		response.Error(r, err)
	}
	response.Success(r, resp)
}
