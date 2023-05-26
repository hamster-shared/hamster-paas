package models

import "time"

type Subscription struct {
	Id                  uint      `json:"id"`
	ChainSubscriptionId uint      `gorm:"column:chain_subscription_id" json:"subscriptionId"`
	Name                string    `json:"name"`
	Created             time.Time `json:"created"`
	Chain               string    `json:"chain"`
	Network             string    `json:"network"`
	Consumers           int8      `json:"consumers"`
	UserId              uint64    `gorm:"column:user_id" json:"userId"`
	Admin               string    `json:"admin"`
	Balance             float64   `json:"balance"`
	TransactionTx       string    `gorm:"column:transaction_tx" json:"transactionTx"`
	Status              string    `json:"status"`
	ErrorMessage        string    `json:"errorMessage"`
}

func (m Subscription) TableName() string {
	return "t_cl_subscription"
}
