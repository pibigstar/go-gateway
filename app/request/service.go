package request

type ServiceInfoListReq struct {
	Content string    `json:"content"`
	Page    *Paginate `json:"page"`
}

type Paginate struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
