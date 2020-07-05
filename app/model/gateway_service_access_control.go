package model

var MGatewayServiceAccessControlModel GatewayServiceAccessControlModel

type GatewayServiceAccessControlModel struct {
	Id                uint64 `gorm:"column:id;primary_key;"     json:"id"`                  // 自增主键
	ServiceId         uint64 `gorm:"column:service_id"          json:"service_id"`          // 服务id
	OpenAuth          int    `gorm:"column:open_auth"           json:"open_auth"`           // 是否开启权限 1=开启
	BlackList         string `gorm:"column:black_list"          json:"black_list"`          // 黑名单ip
	WhiteList         string `gorm:"column:white_list"          json:"white_list"`          // 白名单ip
	WhiteHostName     string `gorm:"column:white_host_name"     json:"white_host_name"`     // 白名单主机
	ClientipFlowLimit int    `gorm:"column:clientip_flow_limit" json:"clientip_flow_limit"` // 客户端ip限流
	ServiceFlowLimit  int    `gorm:"column:service_flow_limit"  json:"service_flow_limit"`  // 服务端限流
}

func (*GatewayServiceAccessControlModel) TableName() string {
	return "gateway_service_access_control"
}
