package admin

import (
	"github.com/gogf/gf/net/ghttp"
	"github/pibigstar/go-gateway/app/model"
	"github/pibigstar/go-gateway/app/request"
	"github/pibigstar/go-gateway/app/response"
	"github/pibigstar/go-gateway/utils/token"
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

	resp, err := model.MAdminModel.AdminLogin(req)
	if err != nil {
		response.Error(r, err)
	}
	adminInfo := &response.AdminInfo{
		Id:       resp.Id,
		UserName: resp.UserName,
	}

	// 生成token
	response.Success(r, token.GenJwtToken(adminInfo))
}

// Login godoc
// @Summary 登陆接口
// @Description 管理员登陆接口
// @Tags 管理员接口
// @ID /admin/info
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=response.Response} "success"
// @Router /admin/info [get]
func Info(r *ghttp.Request) {
	userInfo, err := token.GetUserInfoFromCookie(r)
	if err != nil {
		response.Error(r, err)
		return
	}

	resp, err := model.MAdminModel.AdminInfo(userInfo.Id)
	if err != nil {
		response.Error(r, err)
	}
	adminInfo := &response.AdminInfo{
		Id:       resp.Id,
		UserName: resp.UserName,
	}
	// 生成token
	response.Success(r, adminInfo)
}
