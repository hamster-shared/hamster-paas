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

type ChainLinkRequestExecParam struct {
	SubscriptionId  int64  `json:"subscriptionId"`
	ConsumerAddress string `json:"consumerAddress"`
	Secretsloction  int8   `json:"secretsloction"`
	SecretUrl       string `json:"secretUrl"`
	Args            string `json:"args"`
	TransactionTx   string `json:"transactionTx"`
	RequestName     string `json:"requestName"`
	RequestId       string `json:"requestId"`
}
