package service

import (
	"gorm.io/gorm"
)

type RequestService struct {
	db *gorm.DB
}

func NewRequestService(db *gorm.DB) *RequestService {
	return &RequestService{
		db: db,
	}
}

func (r *RequestService) RequestList() {

}
