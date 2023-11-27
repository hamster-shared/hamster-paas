package service

import (
	"hamster-paas/pkg/models"
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

func (c *ChainLinkDashboardService) GetDashboardAll() *DashboardAll {
	all := &DashboardAll{}
	all.Rpc = c.GetDashboardRpc()
	all.Oracle = c.getMyDashboardOracle()
	all.Storage = c.getMyDashboardStorage()
	all.Graph = c.getMyDashboardGraph()
	all.Zkp = c.getMyDashboardZkp()
	all.Others = c.getMyDashboardOthers()
	return all
}

func (c *ChainLinkDashboardService) GetDashboardOracle() {

}

func (c *ChainLinkDashboardService) GetDashboardRpc() []models.RpcChain {
	s := NewRpcService(c.db, nil)
	chains, err := s.GetChains()
	if err != nil {
		logger.Errorf("GetDashboardRpc error: %s", err)
		return []models.RpcChain{}
	}
	return chains
}

func (c *ChainLinkDashboardService) getMyDashboardOracle() *DashboardOracle {

	return &DashboardOracle{}
}
func (c *ChainLinkDashboardService) getMyDashboardStorage() *DashboardStorage {

	return &DashboardStorage{}
}
func (c *ChainLinkDashboardService) getMyDashboardGraph() *DashboardGraph {

	return &DashboardGraph{}
}

func (c *ChainLinkDashboardService) getMyDashboardZkp() *DashboardZkp {
	return &DashboardZkp{}
}
func (c *ChainLinkDashboardService) getMyDashboardOthers() *DashboardOthers {
	return &DashboardOthers{}
}

func (c *ChainLinkDashboardService) getRpcOveriew() *DashboardOthers {
	return &DashboardOthers{}
}
