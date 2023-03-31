package vo

import "time"

type ChainLinkRequest struct {
	Name   string `json:"name"`
	Script string `json:"script"`
	UserId uint64 `json:"userId"`
}
type ChainLinkRequestParam struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}

type ChainLinkConsumers struct {
	Address    string    `json:"address"`
	Network    string    `json:"network"`
	DeployTime time.Time `json:"deploy_time"`
}
