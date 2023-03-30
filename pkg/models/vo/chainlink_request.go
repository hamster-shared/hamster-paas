package vo

type ChainLinkRequest struct {
	Name   string `json:"name"`
	Script string `json:"script"`
	UserId uint64 `json:"userId"`
}
type ChainLinkRequestParam struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}
