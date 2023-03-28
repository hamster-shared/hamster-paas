package db

type User struct {
	UserID  int64  `gorm:"column:user_id" json:"user_id"`
	Type    string `json:"type"`
	Chain   string `json:"chain"`
	Network string `json:"network"`
}

func (m User) TableName() string {
	return "t_user_middleware"
}
