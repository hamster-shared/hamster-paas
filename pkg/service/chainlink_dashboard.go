package service

import (
	"hamster-paas/pkg/rpc/aline"

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
	Rpc     *DashboardRpc     `json:"rpc"`
	Oracle  *DashboardOracle  `json:"oracle"`
	Storage *DashboardStorage `json:"storage"`
	Graph   *DashboardGraph   `json:"graph"`
	Zkp     *DashboardZkp     `json:"zkp"`
	Others  *DashboardOthers  `json:"others"`
}
type DashboardRpc struct{}
type DashboardOracle struct{}
type DashboardStorage struct{}
type DashboardGraph struct{}
type DashboardZkp struct{}
type DashboardOthers struct{}

func (c *ChainLinkDashboardService) GetDashboardAll(user aline.User) *DashboardAll {
	all := &DashboardAll{}
	all.Rpc = c.getMyDashboardRpc(user)
	all.Oracle = c.getMyDashboardOracle(user)
	all.Storage = c.getMyDashboardStorage(user)
	all.Graph = c.getMyDashboardGraph(user)
	all.Zkp = c.getMyDashboardZkp(user)
	all.Others = c.getMyDashboardOthers(user)
	return all
}

func (c *ChainLinkDashboardService) GetDashboardOracle(user aline.User) *DashboardOracle {

	return &DashboardOracle{}
}

func (c *ChainLinkDashboardService) GetDashboardRpc(user aline.User) *DashboardRpc {

	return &DashboardRpc{}
}

func (c *ChainLinkDashboardService) getMyDashboardRpc(user aline.User) *DashboardRpc {

	return &DashboardRpc{}
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
