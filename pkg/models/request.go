package models

import "time"

type RequestTemplate struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Created     time.Time `json:"created"`
	Script      string    `json:"script"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
}

type Request struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Script  string    `json:"script"`
	UserId  uint64    `gorm:"column:user_id" json:"userId"`
}

type RequestExecute struct {
	Id              int64     `json:"id"`
	SubscriptionId  int64     `gorm:"column:subscription_id" json:"subscriptionId"`
	RequestId       string    `json:"requestId"`
	ConsumerAddress string    `gorm:"column:consumer_address" json:"consumerAddress"`
	Secretsloction  int8      `json:"secretsloction"`
	SecretUrl       string    `json:"secretUrl"`
	Args            string    `json:"args"`
	TransactionTx   string    `gorm:"column:transaction_tx" json:"transactionTx"`
	Status          string    `json:"status"`
	UserId          uint64    `gorm:"column:user_id" json:"userId"`
	RequestName     string    `json:"requestName"`
	Amount          float64   `json:"amount"`
	Created         time.Time `json:"created"`
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
