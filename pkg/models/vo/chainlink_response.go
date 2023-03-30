package vo

import "time"

type ChainLinkRequestVo struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Script  string    `json:"script"`
}

type ChainLinkRequestPage struct {
	Data     []ChainLinkRequestVo `json:"data"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"pageSize"`
}

type RequestTemplateVo struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type ChainLinkSubscriptionOverview struct {
	TotalSubscription int     `json:"total_subscription"`
	TotalConsumers    int     `json:"total_consumers"`
	TotalBalance      float64 `json:"total_balance"`
}
