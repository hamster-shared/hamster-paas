package order

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"hamster-paas/pkg/models/node"
	"time"
)

type OrderType uint

const (
	NodeService OrderType = iota
)

type OrderStatus uint

const (
	PaymentPending OrderStatus = iota + 1
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
	UserId             uint               `json:"userId"`
	OrderTime          sql.NullTime       `json:"orderTime"`
	OrderType          OrderType          `json:"orderType"`
	ResourceType       string             `json:"resourceType"`
	Status             OrderStatus        `json:"status"`
	Chain              node.ChainProtocol `json:"chain"`
	Amount             sql.NullString     `gorm:"type:decimal(10,2)" json:"amount"`
	PayAddress         string             `json:"payAddress"`
	ReceiveAddress     string             `json:"receiveAddress"`
	PayTx              string             `json:"payTx"`
	AddressInitBalance sql.NullString     `gorm:"type:decimal(10,2)" json:"addressInitBalance"`
	BuyTime            int                `json:"buyTime"`
	NodeName           string             `json:"nodeName"`
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

// BlackHeight 区块扫描表
type BlackHeight struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	BlackHeight int64  `json:"blackHeight"`
	EventType   string `json:"eventType"`
}

// ReceiptRecords 收款记录表
type ReceiptRecords struct {
	Id             uint            `gorm:"primaryKey" json:"id"`
	BlackHeight    int64           `json:"blackHeight"`
	Amount         decimal.Decimal `json:"amount"`
	PayAddress     string          `json:"payAddress"`
	ReceiveAddress string          `json:"receiveAddress"`
	PayTx          string          `json:"payTx"`
	OrderId        uint            `json:"orderId"`
	PayTime        time.Time       `json:"payTime"`
	PayTimeUTC     time.Time       `json:"payTimeUTC"`
}
