package models

import "time"

type Deposit struct {
	Id              int64     `json:"id"`
	Created         time.Time `json:"created"`
	RequestName     string    `gorm:"column:request_name" json:"requestName"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumerAddress"`
	Amount          float64   `json:"amount"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transactionTx"`
	Status          string    `json:"status,omitempty"`
	UserId          uint64    `gorm:"column:user_id" json:"userId"`
}

func (m Deposit) TableName() string {
	return "t_cl_deposit"
}
