package model

var MGatewayServiceTcpRuleModel GatewayServiceTcpRuleModel

type GatewayServiceTcpRuleModel struct {
	Id        uint64 `gorm:"column:id;primary_key;" json:"id"`         // 自增主键
	ServiceId uint64 `gorm:"column:service_id"      json:"service_id"` // 服务id
	Port      int    `gorm:"column:port"            json:"port"`       // 端口号
}

func (*GatewayServiceTcpRuleModel) TableName() string {
	return "gateway_service_tcp_rule"
}

func (m *GatewayServiceTcpRuleModel) GetByServiceId(serviceId uint64) (*GatewayServiceTcpRuleModel, error) {
	var result *GatewayServiceTcpRuleModel
	err := db.Table(m.TableName()).Where("service_id = ?", serviceId).Struct(&result)
	return result, err
}
