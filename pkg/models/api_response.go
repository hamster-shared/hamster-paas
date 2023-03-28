package models

type ApiResponse struct {
	Code       int        `json:"code"`
	Message    string     `json:"message"`
	Data       any        `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}
