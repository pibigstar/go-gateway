package admin

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/golang/glog"
	"github/pibigstar/go-gateway/app/consts"
	"github/pibigstar/go-gateway/app/model"
	"github/pibigstar/go-gateway/app/request"
	"github/pibigstar/go-gateway/app/response"
	"github/pibigstar/go-gateway/utils/token"
)

// Login godoc
// @Summary 管理员登陆
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
	t := token.GenJwtToken(adminInfo)
	err = r.Session.Set(consts.UserTokenSessionKey, t)
	if err != nil {
		glog.Error(err)
	}
	response.Success(r, t)
}

// Info godoc
// @Summary 管理员信息
// @Description 管理员信息接口
// @Tags 管理员接口
// @ID /admin/info
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=response.Response} "success"
// @Router /admin/info [get]
func Info(r *ghttp.Request) {
	userInfo, err := token.GetUserInfoFromSession(r)
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
	response.Success(r, adminInfo)
}

// Logout godoc
// @Summary 管理员登出
// @Description 管理员登出接口
// @Tags 管理员接口
// @ID /admin/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=response.Response} "success"
// @Router /admin/logout [get]
func Logout(r *ghttp.Request) {
	err := r.Session.Clear()
	if err != nil {
		response.Error(r, err)
	}
	response.Success(r)
}

// Logout godoc
// @Summary 管理员修改密码
// @Description 管理员修改密码接口
// @Tags 管理员接口
// @ID /admin/changePwd
// @Accept  json
// @Produce  json
// @Param body body request.ChangePasswordReq true "body"
// @Success 200 {object} response.Response{data=response.Response} "success"
// @Router /admin/changePwd [post]
func ChangePwd(r *ghttp.Request) {
	_, err := token.GetUserInfoFromSession(r)
	if err != nil {
		response.Error(r, err)
	}

	var req *request.ChangePasswordReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err)
	}

	if err := model.MAdminModel.ChangePwd(req); err != nil {
		response.Error(r, err)
	}
	response.Success(r)
}
