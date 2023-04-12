package models

import "time"

type Deposit struct {
	Id             int64     `json:"id"`
	SubscriptionId int64     `json:"subscriptionId"`
	Created        time.Time `json:"created"`
	Amount         float64   `json:"amount"`
	TransactionTx  string    `gorm:"column:transaction_tx" json:"transactionTx"`
	Status         string    `json:"status,omitempty"`
	Address        string    `json:"address"`
	UserId         uint64    `gorm:"column:user_id" json:"userId"`
}

func (m Deposit) TableName() string {
	return "t_cl_deposit"
}
