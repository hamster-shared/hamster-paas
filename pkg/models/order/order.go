package order

import (
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
	OrderId string `json:"order_id"`
	// 用户id
	UserId         uint               `json:"user_id"`
	OrderType      OrderType          `json:"order_type"`
	ResourceType   string             `json:"resource_type"`
	Status         OrderStatus        `json:"status"`
	Chain          node.ChainProtocol `json:"chain"`
	Amount         decimal.Decimal    `json:"amount"`
	PayAddress     string             `json:"pay_address"`
	ReceiveAddress string             `json:"receive_address"`
	PayTx          string             `json:"pay_tx"`
}
