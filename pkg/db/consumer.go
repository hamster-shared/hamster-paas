package db

import "time"

type Consumer struct {
	Id              int64     `json:"id"`
	SubscriptionId  int64     `json:"subscriptionId"`
	Created         time.Time `json:"created"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumerAddress"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transactionTx"`
	Status          string    `json:"status"`
	UserId          uint64    `gorm:"column:user_id" json:"userId"`
	CreateTime      time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
}

func (m Consumer) TableName() string {
	return "t_cl_consumer"
}
