package model

var MGatewayServiceGrpcRuleModel GatewayServiceGrpcRuleModel

type GatewayServiceGrpcRuleModel struct {
	Id             uint64 `gorm:"column:id;primary_key;" json:"id"`              // 自增主键
	ServiceId      uint64 `gorm:"column:service_id"      json:"service_id"`      // 服务id
	Port           int    `gorm:"column:port"            json:"port"`            // 端口
	HeaderTransfor string `gorm:"column:header_transfor" json:"header_transfor"` // header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔
}

func (*GatewayServiceGrpcRuleModel) TableName() string {
	return "gateway_service_grpc_rule"
}

func (m *GatewayServiceGrpcRuleModel) GetByServiceId(serviceId uint64) (*GatewayServiceGrpcRuleModel, error) {
	var result *GatewayServiceGrpcRuleModel
	err := db.Table(m.TableName()).Where("service_id = ?", serviceId).Struct(&result)
	return result, err
}
