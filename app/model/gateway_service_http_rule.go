package model

var MGatewayServiceHttpRuleModel GatewayServiceHttpRuleModel

type GatewayServiceHttpRuleModel struct {
	Id             uint64 `gorm:"column:id;primary_key;" json:"id"`              // 自增主键
	ServiceId      uint64 `gorm:"column:service_id"      json:"service_id"`      // 服务id
	RuleType       int    `gorm:"column:rule_type"       json:"rule_type"`       // 匹配类型 0=url前缀url_prefix 1=域名domain
	Rule           string `gorm:"column:rule"            json:"rule"`            // type=domain表示域名，type=url_prefix时表示url前缀
	NeedHttps      int    `gorm:"column:need_https"      json:"need_https"`      // 支持https 1=支持
	NeedStripUri   int    `gorm:"column:need_strip_uri"  json:"need_strip_uri"`  // 启用strip_uri 1=启用
	NeedWebsocket  int    `gorm:"column:need_websocket"  json:"need_websocket"`  // 是否支持websocket 1=支持
	UrlRewrite     string `gorm:"column:url_rewrite"     json:"url_rewrite"`     // url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔
	HeaderTransfor string `gorm:"column:header_transfor" json:"header_transfor"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

func (*GatewayServiceHttpRuleModel) TableName() string {
	return "gateway_service_http_rule"
}
