package models

import "time"

type User struct {
	UserId  int64     `gorm:"column:user_id" json:"userId"`
	Type    string    `json:"type"`
	Chain   string    `json:"chain"`
	Network string    `json:"network"`
	Created time.Time `json:"created"`
}

func (m User) TableName() string {
	return "t_user_middleware"
}
