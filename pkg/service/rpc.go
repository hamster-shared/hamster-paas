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
	if err != nil {
		return nil, err
	}
	for i := range chains {
		chainType, _ := models.ParseChainType(chains[i].Name)
		networkType, _ := models.ParseNetworkType(chains[i].Network)
		chains[i].Name = chainType.String()
		chains[i].Network = networkType.StringWithSpace()
		chains[i].Fullname = fmt.Sprintf("%s %s", chains[i].Name, chains[i].Network)
	}
	return
}

func (s *RpcService) GetChainsWithUserID(userID string) (chains []models.RpcChain, err error) {
	err = s.db.Model(&models.RpcChain{}).Find(&chains).Error
	if err != nil {
		return nil, err
	}
	var userChains []models.RpcApp
	err = s.db.Model(&models.RpcApp{}).Where("account = ?", userID).Find(&userChains).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	for i := range chains {
		chainType, _ := models.ParseChainType(chains[i].Name)
		networkType, _ := models.ParseNetworkType(chains[i].Network)
		chains[i].Name = chainType.String()
		chains[i].Network = networkType.StringWithSpace()
		chains[i].Fullname = fmt.Sprintf("%s %s", chains[i].Name, chains[i].Network)
		for _, userChain := range userChains {
			if userChain.Chain == chains[i].Name && userChain.Network == chains[i].Network {
				chains[i].UserActive = true
			}
		}
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
	a, err := models.GetRpcAccount(fmt.Sprintf("%d", user.Id))
	if err != nil {
		return nil, err
	}
	return a.GetOverview(network)
}

func (s *RpcService) GetMyNetwork(user aline.User, p *models.Pagination) ([]*models.ApiResponseRpcApp, *models.Pagination, error) {
	a, err := models.GetRpcAccount(fmt.Sprintf("%d", user.Id))
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
		logger.Errorf("GetChains error: %s", err)
		return nil, err
	}
	var chainApps []*models.RpcChainApp
	for _, chain := range chains {
		networkType, _ := models.ParseNetworkType(chain.Network)
		a, err := models.GetRpcAccount(fmt.Sprintf("%d", user.Id))
		if err != nil {
			logger.Errorf("GetRpcAccount error: %s", err)
			return nil, err
		}
		app, err := a.GetAppBaseInfoByChainNetwork(chainType, networkType)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			logger.Errorf("GetAppByChainNetwork error: %s", err)
			return nil, err
		}
		chain.NetworkName = fmt.Sprintf("%s %s", chainType.String(), networkType.StringWithSpace())
		chain.Network = networkType.StringWithSpace()
		chain.Name = chainType.String()
		var chainApp models.RpcChainApp
		chainApp.RpcChain = chain
		chainApp.App = app
		chainApps = append(chainApps, &chainApp)
	}
	var detail models.RpcChainDetail
	detail.Chains = chainApps
	if len(chainApps) > 0 {
		name, _ := models.ParseChainType(chainApps[0].Name)
		detail.Name = name.String()
		detail.Image = chainApps[0].Image
	}
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
	a, err := models.GetRpcAccount(fmt.Sprintf("%d", user.Id))
	if err != nil {
		return nil, nil, err
	}
	return a.GetAppRequestLogs(appKey, p)
}

type ServiceIsActiveResponse struct {
	ServiceType string `json:"serviceType"`
	IsActive    bool   `json:"isActive"`
	ChildList   any    `json:"childList"`
}

func (s *RpcService) IsActive(user aline.User, serviceType string) ServiceIsActiveResponse {
	t := strings.ToLower(serviceType)
	if t != string(models.ServiceTypeRpc) && t != string(models.ServiceTypeOracle) {
		return ServiceIsActiveResponse{ServiceType: serviceType, IsActive: false}
	}
	if t == string(models.ServiceTypeRpc) {
		return s.getActiveRpcServiceResponse(fmt.Sprintf("%d", user.Id))
	} else {
		return s.getActiveOracleServiceResponse(fmt.Sprintf("%d", user.Id))
	}
}

func (s *RpcService) getActiveRpcServiceResponse(userID string) ServiceIsActiveResponse {
	var us models.UserService
	err := s.db.Model(&models.UserService{}).Where("user_id = ? and service_type = ?", userID, string(models.ServiceTypeRpc)).First(&us).Error
	if err != nil {
		return ServiceIsActiveResponse{ServiceType: string(models.ServiceTypeRpc), IsActive: false}
	}
	var rpcApps []models.RpcApp
	err = s.db.Model(&models.RpcApp{}).Where("account = ?", userID).Find(&rpcApps).Error
	if err != nil {
		return ServiceIsActiveResponse{ServiceType: string(models.ServiceTypeRpc), IsActive: us.IsActive}
	}
	return ServiceIsActiveResponse{ServiceType: string(models.ServiceTypeRpc), IsActive: us.IsActive, ChildList: rpcApps}
}

func (s *RpcService) getActiveOracleServiceResponse(userID string) ServiceIsActiveResponse {
	var us models.UserService
	err := s.db.Model(&models.UserService{}).Where("user_id = ? and service_type = ?", userID, string(models.ServiceTypeOracle)).First(&us).Error
	if err != nil {
		return ServiceIsActiveResponse{ServiceType: string(models.ServiceTypeOracle), IsActive: false}
	}
	return ServiceIsActiveResponse{ServiceType: string(models.ServiceTypeOracle), IsActive: us.IsActive, ChildList: []string{"Chainlink Functios"}}
}

func (s *RpcService) ActiveService(user aline.User, serviceType, chain, network string) (string, error) {
	t := strings.ToLower(serviceType)
	if t != string(models.ServiceTypeRpc) && t != string(models.ServiceTypeOracle) {
		return "", fmt.Errorf("service type error, only support rpc and oracle")
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
				return "", err
			}
		}
	}
	if t == "rpc" {
		return s.ActiveServiceRpc(user, chain, network)
	}
	return "ok", nil
}

func (s *RpcService) ActiveServiceRpc(user aline.User, chain, network string) (string, error) {
	account, err := models.GetRpcAccount(fmt.Sprintf("%d", user.Id))
	if err != nil {
		return "", err
	}
	_, err = account.CreateAppByString(fmt.Sprintf("%s:%s", strings.ToLower(chain), strings.ToLower(network)), "", chain, network)
	if err != nil {
		return "", err
	}
	return "ok", nil
}
