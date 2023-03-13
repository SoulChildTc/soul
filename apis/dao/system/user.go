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
	result := global.DB.Where("mobile = ?", mobile).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		} else {
			log.Error(result.Error.Error())
		}
	}
	return &user
}

func (s *User) CreateUser(user *model.User) error {
	result := global.DB.Create(user)
	return result.Error
}
