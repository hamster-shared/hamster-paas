package node

import (
	"gorm.io/gorm"
	modelsNode "hamster-paas/pkg/models/node"
	"hamster-paas/pkg/utils/logger"
)

type ResourceStandardService struct {
	db *gorm.DB
}

func NewResourceStandardService(db *gorm.DB) *ResourceStandardService {
	return &ResourceStandardService{
		db: db,
	}
}

// resource standard info
func (r *ResourceStandardService) QueryResourceStandard(chainProtocol, region string) (modelsNode.RpcNodeResourceStandard, error) {
	var nodeSpec modelsNode.RpcNodeResourceStandard
	err := r.db.Model(modelsNode.RpcNodeResourceStandard{}).Where("chain_protocol = ? and region = ?", chainProtocol, region).First(&nodeSpec).Error
	if err != nil {
		logger.Errorf("query resources standard info failed: %s", err)
		return nodeSpec, err
	}
	return nodeSpec, nil
}
