package system

import (
	"errors"
	"gorm.io/gorm"
	"soul/global"
	log "soul/internal/logger"
	"soul/model"
)

type User struct{}

func (s *User) GetUserByMobile(mobile string) *model.User {
	var user model.User
	if err := global.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		log.Error(err.Error())
		return nil
	}

	return &user
}

func (s *User) CreateUser(user *model.User) error {
	result := global.DB.Create(user)
	return result.Error
}
