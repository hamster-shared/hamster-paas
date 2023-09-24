package node

import (
	"database/sql"
	"hamster-paas/pkg/models/node"
	"hamster-paas/pkg/models/order"
)

type OrderVo struct {
	Id           uint               `gorm:"primaryKey" json:"id"`
	OrderId      string             `json:"orderId"`
	OrderTime    sql.NullTime       `json:"orderTime"`
	OrderType    order.OrderType    `json:"orderType"`
	ResourceType string             `json:"resourceType"`
	Status       order.OrderStatus  `json:"status"`
	Chain        node.ChainProtocol `json:"chain"`
	Amount       string             `json:"amount"`
	BuyTime      int                `json:"buyTime"`
	NodeName     string             `json:"nodeName"`
}

type OrderPage struct {
	Data     []OrderVo `json:"data"`
	Total    int64     `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
}

type PayOrderDetail struct {
	OrderVo
	ReceiveAddress string `json:"receiveAddress"`
}
