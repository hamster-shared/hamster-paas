package db

import "time"

type User struct {
	UserId     int64     `gorm:"column:user_id" json:"userId"`
	Type       string    `json:"type"`
	Chain      string    `json:"chain"`
	Network    string    `json:"network"`
	CreateTime time.Time `gorm:"column:create_time;default:current_timestamp" json:"createTime"`
}

func (m User) TableName() string {
	return "t_user_middleware"
}
