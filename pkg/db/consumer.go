package db

import "time"

type Consumer struct {
	Id              int64     `json:"id"`
	Subscription    int64     `json:"subscription"`
	Created         time.Time `json:"created"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumer_address"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transaction_tx"`
	Status          string    `json:"status"`
}

func (m Consumer) TableName() string {
	return "t_cl_consumer"
}
