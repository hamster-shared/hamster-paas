package node

import (
	"database/sql"
	"github.com/shopspring/decimal"
	modelsNode "hamster-paas/pkg/models/node"
)

type NodeStatisticsVo struct {
	Nodes  int64 `json:"nodes"`
	Synced int   `json:"synced"`
	Halted int   `json:"halted"`
}

type NodeVo struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// 节点名
	Name string `json:"name"`
	// 链体系
	ChainProtocol modelsNode.ChainProtocol `json:"chainProtocol"`
	// 节点状态	#see RPCNodeStatus
	Status modelsNode.RPCNodeStatus `json:"status"`
	// 公网IP
	PublicIp string `json:"publicIp"`
	// 节点服务器所在区域
	Region string `json:"region"`
}

type NodePage struct {
	Data     []NodeVo `json:"data"`
	Total    int64    `json:"total"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
}

type NodeDetail struct {
	Id                uint                     `gorm:"primaryKey" json:"id"`
	Name              string                   `json:"name"`
	ChainProtocol     modelsNode.ChainProtocol `json:"chainProtocol"`
	Status            modelsNode.RPCNodeStatus `json:"status"`
	PublicIp          string                   `json:"publicIp"`
	Region            string                   `json:"region"`
	LaunchTime        sql.NullTime             `json:"launchTime"`
	Resource          string                   `json:"resource"`
	ChainVersion      string                   `json:"chainVersion"`
	NextPaymentDate   sql.NullTime             `json:"nextPaymentDate"`
	PaymentPerMonth   decimal.Decimal          `json:"paymentPerMonth"`
	RemainingSyncTime string                   `json:"remainingSyncTime"`
	CurrentHeight     uint                     `json:"currentHeight"`
	BlockTime         string                   `json:"blockTime"`
	HttpEndpoint      string                   `json:"httpEndpoint"`
	WebsocketEndpoint string                   `json:"websocketEndpoint"`
}
