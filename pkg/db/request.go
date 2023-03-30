package db

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
}

type RequestExecute struct {
	Id              int64  `json:"id"`
	SubscriptionId  int64  `gorm:"column:subscription_id" json:"subscription_id"`
	ConsumerAddress string `gorm:"column:consumer_address" json:"consumer_address"`
	Secretsloction  int8   `json:"secretsloction"`
	Args            string `json:"args"`
	TransactionTx   string `gorm:"column:transaction_tx" json:"transaction_tx"`
	Status          string `json:"status"`
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
