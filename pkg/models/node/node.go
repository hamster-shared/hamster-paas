package node

import (
	"database/sql"
	"github.com/shopspring/decimal"
)

type UserNode struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// 用户id
	UserId uint `json:"user_id"`
	// 节点ID
	NodeId uint `json:"node_id"`
}

type RPCNodeStatus uint

const (
	Initializing RPCNodeStatus = iota
	Launching
	Syncing
	Synced
	Halted //停止
)

type ChainProtocol string

const (
	ETHEREUM ChainProtocol = "ethereum"
	APTOS    ChainProtocol = "aptos"
)

type RPCNode struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// 节点名
	Name string `json:"name"`
	// 用户id
	UserId uint `json:"user_id"`
	// 链体系
	Chain ChainProtocol `json:"chain"`
	// 节点状态	#see RPCNodeStatus
	Status RPCNodeStatus `json:"status"`
	// 公网IP
	PublicIp string `json:"public_ip"`
	// 节点服务器所在区域
	Region string `json:"region"`
	// 节点启动时间
	LaunchTime sql.NullTime `json:"launch_time"`
	//资源规格
	Resource string `json:"resource"`
	// 部署的链版本
	ChainVersion string `json:"chain_version"`
	// 下一次支付时间
	NextPaymentDate sql.NullTime `json:"next_payment_date"`
	//每月支付金额
	PaymentPerMonth decimal.Decimal `json:"payment_per_month"`
	// 同步时间
	RemainingSyncTime string `json:"remaining_sync_time"`
	// 当前区块高度
	CurrentHeight uint `json:"current_height"`
	// 平均出块时间
	BlockTime string `json:"block_time"`
	// http rpc 端口地址
	HttpEndpoint string `json:"http_endpoint"`
	// websocket rpc 端口地址
	WebsocketEndpoint string `json:"websocket_endpoint"`
}

type RPCNodeSpec struct {
	Id uint `gorm:"primaryKey" json:"id"`

	ChainProtocol ChainProtocol   `json:"chain_protocol"`
	CPU           string          `json:"cpu"`
	Memory        string          `json:"memory"`
	Disk          string          `json:"disk"`
	CostPerMonth  decimal.Decimal `json:"cost_per_month"`
}
