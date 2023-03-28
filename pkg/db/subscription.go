package db

import "time"

type Subscription struct {
	Id             uint      `json:"id"`
	SubscriptionId uint      `gorm:"column:subscription_id" json:"subscription_id"`
	Name           string    `json:"name"`
	Created        time.Time `json:"created"`
	Chain          string    `json:"chain"`
	Network        string    `json:"network"`
	Consumers      int8      `json:"consumers"`
	Balance        float64   `json:"balance"`
	UserId         uint64    `gorm:"column:user_id" json:"user_id"`
	Admin          string    `json:"admin"`
	TransactionTx  string    `gorm:"column:transaction_tx" json:"transaction_tx"`
	Status         string    `json:"status"`
}

func (m Subscription) TableName() string {
	return "t_cl_subscription"
}
