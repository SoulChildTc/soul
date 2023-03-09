package httputil

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"soul/utils/tags"
	"strings"
)

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func ParseValidateError(err error, obj any) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors.New("数据格式有误")
	}

	var errResult []string
	for _, e := range errs {
		errStr, err := tags.GetTagValue(obj, e.Field(), e.Tag()+"_err")
		if err != nil {
			// 没有获取到就获取默认的msg
			errStr, _ = tags.GetTagValue(obj, e.Field(), "msg")
		}
		errResult = append(errResult, errStr)

	}
	return errors.New(strings.Join(errResult, ","))
}
