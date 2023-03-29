package models

type ApiResponse struct {
	Code       int        `json:"code"`
	Message    string     `json:"message"`
	Data       any        `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Total int64 `json:"total,omitempty"`
	Page  int   `json:"page,omitempty"`
	Size  int   `json:"size,omitempty"`
}
