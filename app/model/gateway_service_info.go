package model

import (
	"time"
)

var MGatewayServiceInfoModel GatewayServiceInfoModel

type GatewayServiceInfoModel struct {
	Id          uint64     `gorm:"column:id;primary_key;" json:"id"`           // 自增主键
	LoadType    int        `gorm:"column:load_type"       json:"load_type"`    // 负载类型 0=http 1=tcp 2=grpc
	ServiceName string     `gorm:"column:service_name"    json:"service_name"` // 服务名称 6-128 数字字母下划线
	ServiceDesc string     `gorm:"column:service_desc"    json:"service_desc"` // 服务描述
	CreateAt    *time.Time `gorm:"column:create_at"       json:"create_at"`    // 添加时间
	UpdateAt    *time.Time `gorm:"column:update_at"       json:"update_at"`    // 更新时间
	IsDelete    int        `gorm:"column:is_delete"       json:"is_delete"`    // 是否删除 1=删除
}

func (*GatewayServiceInfoModel) TableName() string {
	return "gateway_service_info"
}
