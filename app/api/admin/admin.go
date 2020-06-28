package admin

import (
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/model"
	"github/pibigstar/go-gateway/app/request"
	"github/pibigstar/go-gateway/app/response"
)

// Login godoc
// @Summary 登陆接口
// @Description 管理员登陆接口
// @Tags 管理员接口
// @ID /admin/login
// @Accept  json
// @Produce  json
// @Param body body request.AdminLoginReq true "body"
// @Success 200 {object} response.Response{data=response.Response} "success"
// @Router /admin/login [post]
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
