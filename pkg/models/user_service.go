package models

import "time"

type UserService struct {
	Id          int64       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId      int64       `json:"user_id"`
	CreatedAt   time.Time   `json:"created_at"`
	ServiceType ServiceType `json:"service_type"`
	IsActive    bool        `json:"is_active"`
}

type ServiceType string

const (
	ServiceTypeOracle ServiceType = "oracle"
	ServiceTypeRpc    ServiceType = "rpc"
)
