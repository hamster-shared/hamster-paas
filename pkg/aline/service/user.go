package service

import (
	"gorm.io/gorm"
	"hamster-paas/pkg/aline/entity"
	"hamster-paas/pkg/application"
)

type IUserService interface {
	GetUserByToken(token string) (entity.User, error)
}

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

func (u *UserService) GetUserByToken(token string) (entity.User, error) {
	var user entity.User
	res := u.db.Model(entity.User{}).Where("token = ?", token).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

func (u *UserService) UpdateUser(user entity.User) error {
	return u.db.Save(&user).Error
}

func (u *UserService) GetUserById(id int64) (entity.User, error) {
	var user entity.User
	res := u.db.Model(entity.User{}).Where("id = ?", id).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}
