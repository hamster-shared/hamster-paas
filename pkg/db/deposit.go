package db

import "time"

type Deposit struct {
	Id              int64     `json:"id"`
	Created         time.Time `json:"created"`
	RequestName     string    `gorm:"column:request_name" json:"request_name"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumer_address"`
	Amount          float64   `json:"amount"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transaction_tx"`
	Status          string    `json:"status,omitempty"`
}

func (m Deposit) TableName() string {
	return "t_cl_deposit"
}
