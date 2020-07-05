package model

var MGatewayServiceLoadBalanceModel GatewayServiceLoadBalanceModel

type GatewayServiceLoadBalanceModel struct {
	Id                     uint64 `gorm:"column:id;primary_key;"          json:"id"`                       // 自增主键
	ServiceId              uint64 `gorm:"column:service_id"               json:"service_id"`               // 服务id
	CheckMethod            int    `gorm:"column:check_method"             json:"check_method"`             // 检查方法 0=tcpchk,检测端口是否握手成功
	CheckTimeout           int64  `gorm:"column:check_timeout"            json:"check_timeout"`            // check超时时间,单位s
	CheckInterval          int    `gorm:"column:check_interval"           json:"check_interval"`           // 检查间隔, 单位s
	RoundType              int    `gorm:"column:round_type"               json:"round_type"`               // 轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash
	IpList                 string `gorm:"column:ip_list"                  json:"ip_list"`                  // ip列表
	WeightList             string `gorm:"column:weight_list"              json:"weight_list"`              // 权重列表
	ForbidList             string `gorm:"column:forbid_list"              json:"forbid_list"`              // 禁用ip列表
	UpstreamConnectTimeout int64  `gorm:"column:upstream_connect_timeout" json:"upstream_connect_timeout"` // 建立连接超时, 单位s
	UpstreamHeaderTimeout  int64  `gorm:"column:upstream_header_timeout"  json:"upstream_header_timeout"`  // 获取header超时, 单位s
	UpstreamIdleTimeout    int64  `gorm:"column:upstream_idle_timeout"    json:"upstream_idle_timeout"`    // 链接最大空闲时间, 单位s
	UpstreamMaxIdle        uint64 `gorm:"column:upstream_max_idle"        json:"upstream_max_idle"`        // 最大空闲链接数
}

func (*GatewayServiceLoadBalanceModel) TableName() string {
	return "gateway_service_load_balance"
}
