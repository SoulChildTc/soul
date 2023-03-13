package user

import (
	"github.com/gin-gonic/gin"
	"soul/apis/dto"
	"soul/apis/service"
	"soul/utils/httputil"
)

func Login(c *gin.Context) {
	var u dto.Login
	err := c.ShouldBindJSON(&u)
	if err != nil {
		httputil.Error(c, httputil.ParseValidateError(err, &u).Error())
		return
	}
	res, ok := service.UserService.Login(u)
	if ok {
		httputil.OK(c, res, "登录成功")
	} else {
		httputil.Error(c, res)
	}

}

func Register(c *gin.Context) {
	var u dto.Register
	err := c.ShouldBindJSON(&u)
	if err != nil {
		httputil.Error(c, httputil.ParseValidateError(err, &u).Error())
		return
	}
	msg, ok := service.UserService.Register(u)
	if ok {
		httputil.OK(c, nil, msg)
		return
	}
	httputil.Error(c, msg)
}
