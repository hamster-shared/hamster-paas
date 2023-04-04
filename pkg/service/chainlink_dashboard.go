package service

import (
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/rpc/aline"
	"hamster-paas/pkg/utils/logger"

	"gorm.io/gorm"
)

type ChainLinkDashboardService struct {
	db *gorm.DB
}

func NewChainLinkDashboardService(db *gorm.DB) *ChainLinkDashboardService {
	return &ChainLinkDashboardService{
		db: db,
	}
}

type DashboardAll struct {
	Rpc     []models.RpcChain `json:"rpc"`
	Oracle  *DashboardOracle  `json:"oracle"`
	Storage *DashboardStorage `json:"storage"`
	Graph   *DashboardGraph   `json:"graph"`
	Zkp     *DashboardZkp     `json:"zkp"`
	Others  *DashboardOthers  `json:"others"`
}
type DashboardOracle struct{}
type DashboardStorage struct{}
type DashboardGraph struct{}
type DashboardZkp struct{}
type DashboardOthers struct{}

func (c *ChainLinkDashboardService) GetDashboardAll(user aline.User) *DashboardAll {
	all := &DashboardAll{}
	all.Rpc = c.GetDashboardRpc(user)
	all.Oracle = c.getMyDashboardOracle(user)
	all.Storage = c.getMyDashboardStorage(user)
	all.Graph = c.getMyDashboardGraph(user)
	all.Zkp = c.getMyDashboardZkp(user)
	all.Others = c.getMyDashboardOthers(user)
	return all
}

func (c *ChainLinkDashboardService) GetDashboardOracle(user aline.User) {

}

func (c *ChainLinkDashboardService) GetDashboardRpc(user aline.User) []models.RpcChain {
	s := NewRpcService(c.db)
	chains, err := s.GetChains()
	if err != nil {
		logger.Errorf("GetDashboardRpc error: %s", err)
		return []models.RpcChain{}
	}
	return chains
}

func (c *ChainLinkDashboardService) getMyDashboardOracle(user aline.User) *DashboardOracle {

	return &DashboardOracle{}
}
func (c *ChainLinkDashboardService) getMyDashboardStorage(user aline.User) *DashboardStorage {

	return &DashboardStorage{}
}
func (c *ChainLinkDashboardService) getMyDashboardGraph(user aline.User) *DashboardGraph {

	return &DashboardGraph{}
}

func (c *ChainLinkDashboardService) getMyDashboardZkp(user aline.User) *DashboardZkp {
	return &DashboardZkp{}
}
func (c *ChainLinkDashboardService) getMyDashboardOthers(user aline.User) *DashboardOthers {
	return &DashboardOthers{}
}

func (c *ChainLinkDashboardService) getRpcOveriew(user aline.User) *DashboardOthers {
	return &DashboardOthers{}
}
