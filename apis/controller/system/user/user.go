package user

import (
	"github.com/gin-gonic/gin"
	"soul/apis/dto"
	"soul/apis/service"
	"soul/utils/httputil"
)

// Login
// @description 用户登录
// @tags     	User
// @summary		用户登录
// @accept		json
// @produce		json
// @param		data body dto.Login true "手机号,密码"
// @success		200 object httputil.ResponseBody "成功返回token"
// @router		/system/user/login [post]
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

// Register
// @description 用户注册
// @tags     	User
// @summary		用户注册
// @accept		json
// @produce		json
// @param		data body dto.Register true "用户信息"
// @success		200 object httputil.ResponseBody "成功返回"
// @router		/system/user/register [post]
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
