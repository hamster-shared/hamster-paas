package order

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"hamster-paas/pkg/models/node"
)

type OrderType uint

const (
	NodeService OrderType = iota
)

type OrderStatus uint

const (
	PaymentPending OrderStatus = iota
	Paid
	Cancelled
	RefundPending
	Refund
	Dispute
)

type Order struct {
	Id      uint   `gorm:"primaryKey" json:"id"`
	OrderId string `json:"orderId"`
	// 用户id
	UserId         uint               `json:"userId"`
	OrderTime      sql.NullTime       `json:"orderTime"`
	OrderType      OrderType          `json:"orderType"`
	ResourceType   string             `json:"resourceType"`
	Status         OrderStatus        `json:"status"`
	Chain          node.ChainProtocol `json:"chain"`
	Amount         decimal.Decimal    `json:"amount"`
	PayAddress     string             `json:"payAddress"`
	ReceiveAddress string             `json:"receiveAddress"`
	PayTx          string             `json:"payTx"`

	BuyTime int `json:"buyTime"`
}

type OrderNode struct {
	Id         uint         `gorm:"primaryKey" json:"id"`
	OrderId    uint         `json:"orderId"`
	UserId     uint         `json:"userId"`
	NodeName   string       `json:"nodeName"`
	Resource   string       `json:"resource"`
	Protocol   string       `json:"protocol"`
	Region     string       `json:"region"`
	CreateTime sql.NullTime `json:"createTime"`
}
