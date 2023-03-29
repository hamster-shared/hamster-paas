package db

import "time"

type Consumer struct {
	Id              int64     `json:"id"`
	SubscriptionId  int64     `json:"subscription_id"`
	Created         time.Time `json:"created"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumer_address"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transaction_tx"`
	Status          string    `json:"status"`
}

func (m Consumer) TableName() string {
	return "t_cl_consumer"
}
