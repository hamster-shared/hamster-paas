package service

import (
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"strconv"

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

func (s *RpcService) GetChains() (chains []models.RpcChain, err error) {
	err = s.db.Model(&models.RpcChain{}).Find(&chains).Error
	for i := range chains {
		chainType, err := models.ParseChainType(chains[i].Name)
		if err != nil {
			return nil, err
		}
		chains[i].Name = chainType.String()
	}
	if err != nil {
		return nil, err
	}
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

func (s *RpcService) GetMyNetwork(user aline.User, p *models.Pagination) ([]*models.ApiResponseRpcApp, *models.Pagination, error) {
	a, err := models.GetRpcAccount(user.Token)
	if err != nil {
		return nil, p, err
	}
	return a.GetAppsWithPagination(p)
}

func (s *RpcService) ChainDetail(user aline.User, chain string) (*models.RpcChainDetail, error) {
	chainType, err := models.ParseChainType(chain)
	if err != nil {
		return nil, err
	}
	var chains []models.RpcChain
	err = s.db.Model(&models.RpcChain{}).Where("name = ?", chainType.StringLower()).Find(&chains).Error
	if err != nil {
		return nil, err
	}
	var chainApps []*models.RpcChainApp
	for _, chain := range chains {
		networkType, _ := models.ParseNetworkType(chain.Network)
		a, err := models.GetRpcAccount(user.Token)
		if err != nil {
			logger.Errorf("GetRpcAccount error: %s", err)
			return nil, err
		}
		app, err := a.GetAppByChainNetwork(chainType, networkType)
		if err != nil {
			logger.Errorf("GetAppByChainNetwork error: %s", err)
			return nil, err
		}
		var chainApp models.RpcChainApp
		chainApp.RpcChain = chain
		chainApp.App = app
		chainApps = append(chainApps, &chainApp)
	}
	var detail models.RpcChainDetail
	detail.RpcChainBaseInfo = chainType.BaseInfo()
	detail.Chains = chainApps
	return &detail, nil
}

func (s *RpcService) AppRequestLog(user aline.User, appKey, page, size string) ([]*models.RpcAppRequestLog, *models.Pagination, error) {
	var pageInt, sizeInt int
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, nil, err
	}
	sizeInt, err = strconv.Atoi(size)
	if err != nil {
		return nil, nil, err
	}
	p := models.Pagination{
		Page: pageInt,
		Size: sizeInt,
	}
	a, err := models.GetRpcAccount(user.Token)
	if err != nil {
		return nil, nil, err
	}
	return a.GetAppRequestLogs(appKey, p)
}
