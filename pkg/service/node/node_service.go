package node

import (
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	modelsNode "hamster-paas/pkg/models/node"
	"hamster-paas/pkg/models/vo/node"
	"hamster-paas/pkg/utils/logger"
)

type NodeService struct {
	db *gorm.DB
}

func NewNodeService(db *gorm.DB) *NodeService {
	return &NodeService{
		db: db,
	}
}

// node statistic
func (n *NodeService) NodeStatisticsInfo(userId int) (node.NodeStatisticsVo, error) {
	var statisticsInfo node.NodeStatisticsVo
	results := []struct {
		Status int
		Count  int
	}{}
	err := n.db.Model(modelsNode.RPCNode{}).Select("status, count(*) as count").Where("user_id = ? and status in (?,?)", userId, modelsNode.Synced, modelsNode.Halted).
		Group("status").
		Scan(&results).
		Error
	if err != nil {
		logger.Errorf("query node statistics info failed: %s", err)
		return statisticsInfo, err
	}
	if len(results) > 0 {
		for _, result := range results {
			if result.Status == int(modelsNode.Synced) {
				statisticsInfo.Synced = result.Count
			} else {
				statisticsInfo.Halted = result.Count
			}
		}
		statisticsInfo.Nodes = int64(statisticsInfo.Synced + statisticsInfo.Halted)
	}
	return statisticsInfo, nil
}

// node list
func (n *NodeService) NodeList(userId, page, size int) (node.NodePage, error) {
	var nodes []modelsNode.RPCNode
	var data []node.NodeVo
	var total int64
	var nodePage node.NodePage
	err := n.db.Model(modelsNode.RPCNode{}).Where("user_id = ?", userId).Offset((page - 1) * size).Limit(size).Find(&nodes).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		logger.Errorf("query node  list failed: %s", err)
		return nodePage, err
	}
	copier.Copy(&data, &nodes)
	nodePage.Data = data
	nodePage.Page = page
	nodePage.PageSize = size
	nodePage.Total = total
	return nodePage, nil
}

// node detail
func (n *NodeService) NodeDetail(nodeId int) (node.NodeDetail, error) {
	var nodeData modelsNode.RPCNode
	var detailData node.NodeDetail
	err := n.db.Model(modelsNode.RPCNode{}).Where("id = ?", nodeId).First(&nodeData).Error
	if err != nil {
		logger.Errorf("query node  detail failed: %s", err)
		return detailData, err
	}
	copier.Copy(&detailData, &nodeData)
	return detailData, nil
}

func (n *NodeService) SaveNode(userId int, nodeData node.SaveNodeParam) error {
	var node modelsNode.RPCNode
	err := n.db.Where("name=? and user_id=?", nodeData.Name, userId).First(&node).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			copier.Copy(&node, &nodeData)
			err = n.db.Create(&node).Error
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return fmt.Errorf("node name: %s is aready exists", nodeData.Name)
}
