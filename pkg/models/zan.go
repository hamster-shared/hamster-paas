package models

import "time"

type ZanUser struct {
	ID          uint      `json:"id"`
	UserId      string    `json:"user_id"`
	AccessToken string    `json:"access_token"`
	Created     time.Time `json:"created"`
}

func (z *ZanUser) TableName() string {
	return "t_zan_user"
}
