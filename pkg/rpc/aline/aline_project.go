package aline

import (
	"gorm.io/gorm"
	"hamster-paas/pkg/models"
	"hamster-paas/pkg/models/vo"
)

type ProjectService struct {
	db *gorm.DB
}

func NewAlineProjectService(db *gorm.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

func (p *ProjectService) GetProjectByUserId(userId uint) []*vo.AlineProjectIDAndName {
	var projectList []*vo.AlineProjectIDAndName
	p.db.Model(models.Project{}).Where("user_id = ?", userId).Find(&projectList)
	return projectList
}
