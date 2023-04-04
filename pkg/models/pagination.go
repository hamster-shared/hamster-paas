package models

type Pagination struct {
	Total int64 `json:"total,omitempty"`
	Page  int   `json:"page,omitempty"`
	Size  int   `json:"size,omitempty"`
}
