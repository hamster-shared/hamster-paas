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
	TotalSubscription int `json:"total_subscription"`
	TotalConsumers    int `json:"total_consumers"`
}

type ChainLinkSINA struct {
	SubscriptionId int64   `json:"subscription_id"`
	Name           string  `json:"name"`
	Balance        float64 `gorm:"balance" json:"balance"`
}

type ChainLinkSubscriptionPage struct {
	Data     []ChainLinkSubscriptionVo `json:"data"`
	Total    int64                     `json:"total"`
	Page     int                       `json:"page"`
	PageSize int                       `json:"pageSize"`
}

type ChainLinkSubscriptionVo struct {
	Id            uint      `json:"id"`
	Name          string    `json:"name"`
	Created       time.Time `json:"created"`
	Chain         string    `json:"chain"`
	Network       string    `json:"network"`
	Consumers     int8      `json:"consumers"`
	Balance       float64   `json:"balance"`
	TransactionTx string    `json:"transactionTx"`
	Admin         string    `json:"admin"`
	Status        string    `json:"status"`
}

type ChainLinkConsumerPage struct {
	Data     []ChainLinkConsumerVo `json:"data"`
	Total    int64                 `json:"total"`
	Page     int                   `json:"page"`
	PageSize int                   `json:"pageSize"`
}

type ChainLinkConsumerVo struct {
	Id              int64     `json:"id"`
	Created         time.Time `json:"created"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumerAddress"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transactionTx"`
	Status          string    `json:"status"`
}

type ChainLinkDepositPage struct {
	Data     []ChainLinkDepositVo `json:"data"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"pageSize"`
}

type ChainLinkDepositVo struct {
	Id              int64     `json:"id"`
	Created         time.Time `json:"created"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumerAddress"`
	Amount          float64   `json:"amount"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transactionTx"`
	Status          string    `json:"status,omitempty"`
}

type ChainLinkExpensePage struct {
	Data     []ChainLinkExpenseVo `json:"data"`
	Total    int64                `json:"total"`
	Page     int                  `json:"page"`
	PageSize int                  `json:"pageSize"`
}

type ChainLinkExpenseVo struct {
	Id              int64     `json:"id"`
	ConsumerAddress string    `json:"consumerAddress"`
	TransactionTx   string    `json:"transactionTx"`
	Status          string    `json:"status"`
	RequestName     string    `json:"requestName"`
	Amount          float64   `json:"amount"`
	Created         time.Time `json:"created"`
}

type ChainLinkValidSubscriptionVo struct {
	Id            uint      `json:"id"`
	Name          string    `json:"name"`
	Created       time.Time `json:"created"`
	Chain         string    `json:"chain"`
	Network       string    `json:"network"`
	TransactionTx string    `json:"transactionTx"`
	Admin         string    `json:"admin"`
	Status        string    `json:"status"`
}
