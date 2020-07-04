package request

type AdminLoginReq struct {
	UserName string `p:"username" v:"required|length:5,30#请输入账号|账号长度为:min到:max位"` // 用户名
	Password string `p:"password" v:"required|length:5,30#请输入密码|密码长度为:min到:max位"` // 密码
}

type ChangePasswordReq struct {
	UserName    string `p:"username" v:"required|length:5,30#请输入账号|账号长度为:min到:max位"`
	Password    string `p:"password" v:"required|length:5,30#请输入密码|密码长度为:min到:max位"`    //密码
	NewPassword string `p:"newPassword" v:"required|length:5,30#请输入密码|密码长度为:min到:max位"` //密码
}
