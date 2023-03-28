package models

type ApiRequestCreateApp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Chain       string `json:"chain"`
	Network     string `json:"network"`
	Account     string `json:"account"`
}

type ApiRequestPagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
