package system

import (
	"soul/apis/dao"
	"soul/apis/dto"
	log "soul/internal/logger"
	"soul/model"
	"soul/utils"
)

type User struct{}

func (s *User) Register(user dto.Register) (string, bool) {
	if user.Username == "admin" {
		return "禁止使用admin注册", false
	}

	existUser := dao.UserDAO.GetUserByMobile(user.Mobile)
	if existUser != nil {
		return "手机号已存在", false
	}

	// 创建新用户
	newUser := &model.User{
		Username: user.Username,
		Mobile:   user.Mobile,
		Password: utils.PasswdMd5Digest(user.Password),
		Email:    user.Email,
	}

	err := dao.UserDAO.CreateUser(newUser)
	if err != nil {
		return err.Error(), false
	}

	return "注册成功", true
}

func (s *User) Login(user dto.Login) (string, bool) {
	log.Info("test service")
	existUser := dao.UserDAO.GetUserByMobile(user.Mobile)
	if existUser == nil {
		return "账号不存在", false
	}
	if existUser.Password != utils.PasswdMd5Digest(user.Password) {
		return "手机号或密码错误", false
	}

	token, err := utils.CreateJwtToken(int(existUser.ID.ID), existUser.Username)
	if err != nil {
		log.Error(err.Error())
		return "生成token发生错误", false
	}

	return token, true
}
