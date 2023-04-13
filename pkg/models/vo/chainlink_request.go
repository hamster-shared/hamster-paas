package vo

import "time"

type ChainLinkRequest struct {
	Name        string `json:"name"`
	Script      string `json:"script"`
	UserId      uint64 `json:"userId"`
	ParamsCount int    `json:"paramsCount"`
}
type ChainLinkRequestParam struct {
	Name        string `json:"name"`
	Script      string `json:"script"`
	ParamsCount int    `json:"paramsCount"`
}

type ChainLinkConsumers struct {
	Address    string    `json:"address"`
	Network    string    `json:"network"`
	DeployTime time.Time `json:"deploy_time"`
}

type ChainLinkRequestExecParam struct {
	Network         string  `json:"network"`
	SubscriptionId  int64   `json:"subscriptionId"`
	ConsumerAddress string  `json:"consumerAddress"`
	Secretsloction  int8    `json:"secretsloction"`
	SecretUrl       string  `json:"secretUrl"`
	Args            string  `json:"args"`
	TransactionTx   string  `json:"transactionTx"`
	RequestName     string  `json:"requestName"`
	RequestId       string  `json:"requestId"`
	Amount          float64 `json:"amount"`
}

type ChainLinkExecParam struct {
	Network   string `json:"network"`
	RequestId string `json:"requestId"`
}

type ChainLinkSubscriptionCreateParam struct {
	Chain          string `json:"chain"`
	Network        string `json:"network"`
	Name           string `json:"name"`
	SubscriptionId int64  `json:"subscriptionId"`
	Admin          string `json:"admin"`
	TransactionTx  string `json:"transactionTx"`
}

// ChainLinkFoundParam address是admin地址
type ChainLinkFoundParam struct {
	Address       string `json:"address"`
	Incr          string `json:"incr"`
	TransactionTx string `json:"transactionTx"`
}

type ChainLinkConsumerCreateParam struct {
	SubscriptionId  int64  `json:"subscriptionId"`
	ConsumerAddress string `json:"consumerAddress"`
	TransactionTx   string `json:"transactionTx"`
}

type ChainLinkFoundUpdateParam struct {
	Id             int64  `json:"id"`
	SubscriptionId int64  `json:"subscriptionId"`
	TransactionTx  string `json:"transactionTx"`
	NewStatus      string `json:"newStatus"`
}

type ChainLinkConsumerUpdateParam struct {
	Id              int64  `json:"id"`
	SubscriptionId  int64  `json:"subscriptionId"`
	ConsumerAddress string `json:"consumerAddress"`
	TransactionTx   string `json:"transactionTx"`
	NewStatus       string `json:"newStatus"`
}

type ChainLinkSubscriptionUpdateParam struct {
	Id                  uint   `json:"id"`
	ChainSubscriptionId uint   `json:"chainSubscriptionId"`
	Chain               string `json:"chain"`
	Network             string `json:"network"`
	TransactionTx       string `json:"transactionTx"`
	NewStatus           string `json:"newStatus"`
}

type ChainLinkHamsterListParam struct {
	Chain   string `json:"chain"`
	NetWork string `json:"network"`
}
