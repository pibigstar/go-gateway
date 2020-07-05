package model

import "time"

var MGatewayAppModel GatewayAppModel

type GatewayAppModel struct {
	Id       uint64     `gorm:"column:id;primary_key;" json:"id"`        // 自增id
	AppId    string     `gorm:"column:app_id"          json:"app_id"`    // 租户id
	Name     string     `gorm:"column:name"            json:"name"`      // 租户名称
	Secret   string     `gorm:"column:secret"          json:"secret"`    // 密钥
	WhiteIps string     `gorm:"column:white_ips"       json:"white_ips"` // ip白名单，支持前缀匹配
	Qpd      int64      `gorm:"column:qpd"             json:"qpd"`       // 日请求量限制
	Qps      int64      `gorm:"column:qps"             json:"qps"`       // 每秒请求量限制
	CreateAt *time.Time `gorm:"column:create_at"       json:"create_at"` // 添加时间
	UpdateAt *time.Time `gorm:"column:update_at"       json:"update_at"` // 更新时间
	IsDelete int        `gorm:"column:is_delete"       json:"is_delete"` // 是否删除 1=删除
}

func (*GatewayAppModel) TableName() string {
	return "gateway_app"
}
