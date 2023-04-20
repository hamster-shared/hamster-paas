package service

import (
	"hamster-paas/pkg/models"

	"gorm.io/gorm"
)

type MiddleWareService struct {
	db *gorm.DB
}

func NewMiddleWareService(db *gorm.DB) *MiddleWareService {
	return &MiddleWareService{
		db: db,
	}
}

func (s *MiddleWareService) MiddleWareRpc(userID string) ([]*models.MiddleWareRpc, error) {
	rpcService := NewRpcService(s.db)
	chains, err := rpcService.GetChainsWithUserID(userID)
	if err != nil {
		return nil, err
	}
	var result []*models.MiddleWareRpc
	for i := range chains {
		result = append(result, &models.MiddleWareRpc{
			RpcChain: chains[i],
		})
	}
	return result, nil
}
