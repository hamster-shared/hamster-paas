package service

import (
	"fmt"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"
	"strconv"
	"strings"

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

func (s *RpcService) Overview(user aline.User, network string) (*models.ApiResponseOverview, error) {
	a, err := models.GetRpcAccount(user.Token)
	if err != nil {
		return nil, err
	}
	return a.GetOverview(network)
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
		app, err := a.GetAppBaseInfoByChainNetwork(chainType, networkType)
		if err != nil {
			logger.Errorf("GetAppByChainNetwork error: %s", err)
			return nil, err
		}
		chain.NetworkName = fmt.Sprintf("%s %s", chainType.String(), networkType.StringWithSpace())
		var chainApp models.RpcChainApp
		chainApp.RpcChain = chain
		chainApp.App = app
		chainApps = append(chainApps, &chainApp)
	}
	var detail models.RpcChainDetail
	// detail.RpcChainBaseInfo = chainType.BaseInfo()
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

func (s *RpcService) IsActive(user aline.User, serviceType string) bool {
	t := strings.ToLower(serviceType)
	if t != string(models.ServiceTypeRpc) && t != string(models.ServiceTypeOracle) {
		return false
	}
	var us models.UserService
	err := s.db.Model(&models.UserService{}).Where("user_id = ? and service_type = ?", user.Id, strings.ToLower(serviceType)).First(&us).Error
	if err != nil {
		return false
	}
	return us.IsActive
}

func (s *RpcService) ActiveService(user aline.User, serviceType string) string {
	t := strings.ToLower(serviceType)
	if t != string(models.ServiceTypeRpc) && t != string(models.ServiceTypeOracle) {
		return "service type error, only support rpc and oracle"
	}
	var us models.UserService
	err := s.db.Model(&models.UserService{}).Where("user_id = ? and service_type = ?", user.Id, t).First(&us).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			us.UserId = int64(user.Id)
			us.ServiceType = models.ServiceType(t)
			us.IsActive = true
			err = s.db.Model(&models.UserService{}).Create(&us).Error
			if err != nil {
				return err.Error()
			}
			return "ok"
		}
		return err.Error()
	}
	return "service already active"
}
