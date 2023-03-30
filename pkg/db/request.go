package db

import "time"

type RequestTemplate struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Created     time.Time `json:"created"`
	Script      string    `json:"script"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	CreateTime  time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
}

type Request struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Created    time.Time `json:"created"`
	Script     string    `json:"script"`
	UserId     uint64    `gorm:"column:user_id" json:"userId"`
	CreateTime time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
}

type RequestExecute struct {
	Id              int64     `json:"id"`
	SubscriptionId  int64     `gorm:"column:subscription_id" json:"subscriptionId"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumerAddress"`
	Secretsloction  int8      `json:"secretsloction"`
	Args            string    `json:"args"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transactionTx"`
	Status          string    `json:"status"`
	UserId          uint64    `gorm:"column:user_id" json:"userId"`
	CreateTime      time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
}

func (m RequestTemplate) TableName() string {
	return "t_cl_request_template"
}

func (m Request) TableName() string {
	return "t_cl_request"
}

func (m RequestExecute) TableName() string {
	return "t_cl_request_execute"
}
