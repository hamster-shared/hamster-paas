package node

import (
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
	// 用户id
	UserId uint `json:"userId"`
	// 链体系
	Chain modelsNode.ChainProtocol `json:"chain"`
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
