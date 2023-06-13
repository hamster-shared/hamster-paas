package node

import (
	"database/sql"
	"github.com/shopspring/decimal"
	modelsNode "hamster-paas/pkg/models/node"
)

type NodeLaunchParam struct {
	Name     string                   `json:"name"`
	Chain    modelsNode.ChainProtocol `json:"chain"`
	Region   string                   `json:"region"`
	Resource string                   `json:"resource"`
	// 下一次支付时间
	NextPaymentDate sql.NullTime `json:"nextPaymentDate"`
	//每月支付金额
	PaymentPerMonth decimal.Decimal `json:"paymentPerMonth"`
}

type SaveNodeParam struct {
	Name            string                   `json:"name"`
	UserId          uint                     `json:"userId"`
	Chain           modelsNode.ChainProtocol `json:"chain"`
	Status          modelsNode.RPCNodeStatus `json:"status"`
	Region          string                   `json:"region"`
	Resource        string                   `json:"resource"`
	NextPaymentDate sql.NullTime             `json:"nextPaymentDate"`
	PaymentPerMonth decimal.Decimal          `json:"paymentPerMonth"`
}
