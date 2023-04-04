package service

import (
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"

	"gorm.io/gorm"
)

type RpcService struct {
	db *gorm.DB
}

func NewRpcService(db *gorm.DB) *RpcService {
	return &RpcService{
		db: db,
	}
}

func (s *RpcService) GetChains(user aline.User) (chains []models.RpcChain, err error) {
	err = s.db.Model(&models.RpcChain{}).Find(&chains).Error
	if err != nil {
		return nil, err
	}
	// u, err := models.GetRpcAccount(user.Token)
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println(user)
	// for i := range chains {
	// 	if chains[i].HttpAddress != "" {
	// 		chains[i].HttpAddress += user.AvatarUrl
	// 	}
	// }
	return
}

func (s *RpcService) GetNetworks(chain string) ([]string, error) {
	chainType, err := models.ParseChainType(chain)
	if err != nil {
		return nil, err
	}
	var chains []models.RpcChain
	err = s.db.Model(&models.RpcChain{}).Where("name = ?", chainType.String()).Find(&chains).Error
	if err != nil {
		return nil, err
	}
	var networks []string
	for _, chain := range chains {
		networks = append(networks, chain.Network)
	}
	return networks, nil
}

func (s *RpcService) Overview(user aline.User) ([]*models.ApiResponseRpcApp, error) {
	a, err := models.GetRpcAccount(user.Token)
	if err != nil {
		return nil, err
	}
	return a.GetApps()
}
