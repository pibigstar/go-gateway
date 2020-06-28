package model

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github/pibigstar/go-gateway/app/request"
	"github/pibigstar/go-gateway/utils"
)

var MAdminModel AdminModel

type AdminModel struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
	IsDelete bool   `json:"isDelete"`
}

func (*AdminModel) Table() string {
	return "gateway_admin"
}

func (a *AdminModel) AdminLogin(ctx context.Context, req *request.AdminLoginReq) (*AdminModel, error) {
	record, err := g.DB().Table(a.Table()).
		Where("user_name = ?", req.UserName).
		Where("is_delete = 0").One()
	if err != nil {
		return nil, err
	}
	var adminModel *AdminModel
	err = record.Struct(&adminModel)
	if err != nil {
		return nil, err
	}
	pwd := utils.GenSaltPassword(adminModel.Salt, req.Password)
	if adminModel.Password != pwd {
		return nil, errors.New("密码错误")
	}
	return adminModel, err
}
