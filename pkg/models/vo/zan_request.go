package vo

type ExchangeAccessTokenVo struct {
	// authCode 交换accessToken的授权码
	AuthCode string `json:"authCode"`
}

type PageResp[T interface{}] struct {
	// 总条⽬数
	Total int64 `json:"total"`
	// 总页数
	PageCount int `json:"pageCount"`
	// 当前页
	Page int `json:"page"`
	// 每页条目数
	PageSize int `json:"pageSize"`
	// 条⽬列表
	Data []T `json:"data"`
}

func NewPageResp[T interface{}](page int, size int, data []T) PageResp[T] {
	return PageResp[T]{
		Total:     int64(len(data)),
		PageCount: len(data) / size,
		Page:      page,
		PageSize:  size,
		Data:      data,
	}
}
