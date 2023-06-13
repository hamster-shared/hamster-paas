package node

import (
	"database/sql"
	"github.com/shopspring/decimal"
)

type UserNode struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// 用户id
	UserId uint `json:"userId"`
	// 节点ID
	NodeId uint `json:"nodeId"`
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
	ETHEREUM  ChainProtocol = "Ethereum"
	APTOS     ChainProtocol = "Aptos"
	SUI       ChainProtocol = "Sui"
	NEAR      ChainProtocol = "Near"
	STARK_NET ChainProtocol = "StarkNet"
	OPTIMISM  ChainProtocol = "Optimism"
	AVALANCHE ChainProtocol = "Avalanche"
)

type RPCNode struct {
	Id uint `gorm:"primaryKey" json:"id"`
	// 节点名
	Name string `json:"name"`
	// 用户id
	UserId uint `json:"userId"`
	// 链体系
	Chain ChainProtocol `json:"chain"`
	// 节点状态	#see RPCNodeStatus
	Status RPCNodeStatus `json:"status"`
	// 公网IP
	PublicIp string `json:"publicIp"`
	// 节点服务器所在区域
	Region string `json:"region"`
	// 节点启动时间
	LaunchTime sql.NullTime `json:"launchTime"`
	//资源规格
	Resource string `json:"resource"`
	// 部署的链版本
	ChainVersion string `json:"chainVersion"`
	// 下一次支付时间
	NextPaymentDate sql.NullTime `json:"nextPaymentDate"`
	//每月支付金额
	PaymentPerMonth decimal.Decimal `json:"paymentPerMonth"`
	// 同步时间
	RemainingSyncTime string `json:"remainingSyncTime"`
	// 当前区块高度
	CurrentHeight uint `json:"currentHeight"`
	// 平均出块时间
	BlockTime string `json:"blockTime"`
	// http rpc 端口地址
	HttpEndpoint string `json:"httpEndpoint"`
	// websocket rpc 端口地址
	WebsocketEndpoint string `json:"websocketEndpoint"`
}

type RpcNodeResourceStandard struct {
	Id            uint            `gorm:"primaryKey" json:"id"`
	ChainProtocol ChainProtocol   `json:"chainProtocol"`
	CPU           string          `json:"cpu"`
	Memory        string          `json:"memory"`
	Disk          string          `json:"disk"`
	Region        string          `json:"region"`
	CostPerMonth  decimal.Decimal `json:"costPerMonth"`
}

type UserChargeAccount struct {
	UserId  int    `json:"userId"`
	Seed    string `json:"seed"`
	Address string `json:"address"`
}
