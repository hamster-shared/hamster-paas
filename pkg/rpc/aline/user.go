package aline

import (
	"gorm.io/gorm"
	"hamster-paas/pkg/application"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	user := UserService{}
	alineDb, err := application.GetBean[*gorm.DB]("alineDb")
	if err != nil {
		return nil
	}
	user.db = alineDb
	return &user
}

func (u *UserService) GetUserByToken(token string) (User, error) {
	var user User
	res := u.db.Model(User{}).Where("token = ?", token).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (u *UserService) UpdateUser(user User) error {
	return u.db.Save(&user).Error
}

func (u *UserService) GetUserById(id int64) (User, error) {
	var user User
	res := u.db.Model(User{}).Where("id = ?", id).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
