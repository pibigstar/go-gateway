package request

type ServiceInfoListReq struct {
	Content string    `json:"content"`
	Page    *Paginate `json:"page"`
}

type Paginate struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type ServiceDetailReq struct {
	Id uint64 `p:"id" v:"required#请输入服务ID"`
}
