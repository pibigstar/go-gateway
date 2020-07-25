package response

import (
	"github.com/gogf/gf/os/gtime"
)

type ServiceInfo struct {
	Id          uint64     `json:"id"`           // 自增主键
	LoadType    int        `json:"load_type"`    // 负载类型 0=http 1=tcp 2=grpc
	ServiceName string     `json:"service_name"` // 服务名称 6-128 数字字母下划线
	ServiceDesc string     `json:"service_desc"` // 服务描述
	CreateAt    gtime.Time `json:"create_at"`    // 添加时间
	UpdateAt    gtime.Time `json:"update_at"`    // 更新时间
}

type ServiceInfoListResp struct {
	List  []*ServiceInfo `json:"list"`
	Total int            `json:"total"`
}

type ServiceDetail struct {
	Id          uint64     `json:"id"`           // 自增主键
	LoadType    int        `json:"load_type"`    // 负载类型 0=http 1=tcp 2=grpc
	ServiceName string     `json:"service_name"` // 服务名称 6-128 数字字母下划线
	ServiceDesc string     `json:"service_desc"` // 服务描述
	HTTP        string     `json:"http"`         // http地址
	GRPC        string     `json:"grpc"`         // grpc地址
	TCP         string     `json:"tcp"`          // tcp地址描述
	CreateAt    gtime.Time `json:"create_at"`    // 添加时间
	UpdateAt    gtime.Time `json:"update_at"`    // 更新时间
}
