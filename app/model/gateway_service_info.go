package model

import (
	"github.com/gogf/gf/os/gtime"
	"github/pibigstar/go-gateway/app/request"
)

var MServiceInfoModel ServiceInfoModel

type ServiceInfoModel struct {
	Id          uint64     `json:"id"`           // 自增主键
	LoadType    int        `json:"load_type"`    // 负载类型 0=http 1=tcp 2=grpc
	ServiceName string     `json:"service_name"` // 服务名称 6-128 数字字母下划线
	ServiceDesc string     `json:"service_desc"` // 服务描述
	CreateAt    gtime.Time `json:"create_at"`    // 添加时间
	UpdateAt    gtime.Time `json:"update_at"`    // 更新时间
	IsDelete    int        `json:"is_delete"`    // 是否删除 1=删除
}

func (*ServiceInfoModel) TableName() string {
	return "gateway_service_info"
}

func (m *ServiceInfoModel) Get(id uint64) (*ServiceInfoModel, error) {
	one, err := db.Table(m.TableName()).Where("id = ?", id).FindOne()
	if err != nil {
		return nil, err
	}
	var info *ServiceInfoModel
	err = one.Struct(&info)
	return info, err
}

func (m *ServiceInfoModel) PageList(req *request.ServiceInfoListReq) ([]*ServiceInfoModel, int, error) {
	db := db.Table(m.TableName())
	if req.Content != "" {
		db = db.Where("service_name like ?", req.Content+"%").
			Or("service_desc like ?", req.Content+"%")
	}
	total, err := db.Clone().Count()
	if err != nil || total == 0 {
		return nil, total, err
	}

	if req.Page != nil {
		db = db.Page(req.Page.Page, req.Page.Size)
	}
	var results []*ServiceInfoModel
	record, err := db.FindAll()
	if err != nil {
		return nil, 0, err
	}

	err = record.Structs(&results)
	if err != nil {
		return nil, 0, err
	}

	return results, total, err

}
