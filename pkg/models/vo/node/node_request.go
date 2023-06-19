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

type UpdateNodeParam struct {
	ID                uint                     `json:"id"`
	Name              string                   `json:"name"`
	PublicIp          string                   `json:"publicIp"`
	ChainVersion      modelsNode.ChainProtocol `json:"chainVersion"`
	Status            modelsNode.RPCNodeStatus `json:"status"`
	RemainingSyncTime string                   `json:"remainingSyncTime"`
	HttpEndpoint      string                   `json:"httpEndpoint"`
	WebsocketEndpoint string                   `json:"websocketEndpoint"`
	VerifyIdentity    string                   `json:"verifyIdentity"`
}
