package aline

import (
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
	"strings"
)

type ProjectService struct {
	db *gorm.DB
}

func NewAlineProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

func (p *ProjectService) GetProjectByUserId(userId uint, network string) []*vo.AlineProjectIDAndName {
	var projectList []*vo.AlineProjectIDAndName
	var projectListRet []*vo.AlineProjectIDAndName
	p.db.Model(models.Project{}).Where("user_id = ? AND label_display = ?", userId, "Chainlink").Find(&projectList)
	// TODO 仅显示拥有对应chain network的合约的project id
	split := strings.Split(network, "/")
	if len(split) > 1 {
		network = split[1]
	}
	for _, v := range projectList {
		var c int64
		p.db.Model(models.ContractDeploy{}).Where("project_id = ? AND network = ?", v.Id, network).Count(&c)
		if c != 0 {
			projectListRet = append(projectListRet, v)
		}
	}
	return projectListRet
}

func (p *ProjectService) GetValidContract(page, size int, projectId string, network string) (*vo.AlineValidContractPage, error) {
	var total int64
	var alineValidContractPage vo.AlineValidContractPage
	var alineValidContractList []models.ContractDeploy
	var alineValidContractVoList []vo.AlineValidContractVo
	tx := p.db.Model(models.ContractDeploy{}).Where("project_id = ? AND network = ?", projectId, network)
	result := tx.Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&alineValidContractList).Offset(-1).Limit(-1).Count(&total)
	if result.Error != nil {
		return &alineValidContractPage, result.Error
	}
	copier.Copy(&alineValidContractVoList, &alineValidContractList)
	for i, v := range alineValidContractVoList {
		n := models.AlineNetworkToChainkLinkNetwork(v.Network)
		alineValidContractVoList[i].Network = n
	}
	alineValidContractPage.Data = alineValidContractVoList
	alineValidContractPage.Total = total
	alineValidContractPage.Page = page
	alineValidContractPage.PageSize = size
	return &alineValidContractPage, nil
}
