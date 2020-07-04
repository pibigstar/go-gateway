package model

import (
	"github.com/gogf/gf/database/gdb"
	"github/pibigstar/go-gateway/app/consts/code"
	"github/pibigstar/go-gateway/app/request"
	"github/pibigstar/go-gateway/utils"
	"github/pibigstar/go-gateway/utils/errx"
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

func (a *AdminModel) AdminLogin(req *request.AdminLoginReq) (*AdminModel, error) {
	record, err := db.Table(a.Table()).Fields("id,user_name,salt,password").
		Where("user_name = ?", req.UserName).
		Where("is_delete = 0").One()
	if err != nil {
		if gdb.ErrNoRows == err {
			return nil, errx.New(code.Error_User_Not_Exist)
		}
		return nil, err
	}

	var adminModel *AdminModel
	err = record.Struct(&adminModel)
	if err != nil {
		return nil, err
	}

	pwd := utils.GenSaltPassword(adminModel.Salt, req.Password)
	if adminModel.Password != pwd {
		return nil, errx.New(code.Error_Password_Error)
	}
	return adminModel, err
}

func (a *AdminModel) AdminInfo(userId int64) (*AdminModel, error) {
	record, err := db.Table(a.Table()).Fields("id,user_name").
		Where("id = ?", userId).
		Where("is_delete = 0").One()
	if err != nil {
		if gdb.ErrNoRows == err {
			return nil, errx.New(code.Error_User_Not_Exist)
		}
		return nil, err
	}

	var adminModel *AdminModel
	err = record.Struct(&adminModel)
	if err != nil {
		return nil, err
	}

	return adminModel, err
}

func (a *AdminModel) ChangePwd(req *request.ChangePasswordReq) error {
	record, err := db.Table(a.Table()).Fields("id,user_name,salt,password").
		Where("user_name = ?", req.UserName).
		Where("is_delete = 0").One()
	if err != nil {
		if gdb.ErrNoRows == err {
			return errx.New(code.Error_User_Not_Exist)
		}
		return err
	}

	var adminModel *AdminModel
	err = record.Struct(&adminModel)
	if err != nil {
		return err
	}

	pwd := utils.GenSaltPassword(adminModel.Salt, req.Password)
	if adminModel.Password != pwd {
		return errx.New(code.Error_Password_Error)
	}
	// 更换密码
	encryptPwd := utils.GenSaltPassword(adminModel.Salt, req.NewPassword)
	_, err = db.Table(a.Table()).
		Data("password", encryptPwd).
		Where("id = ?", adminModel.Id).
		Update()
	return err
}
