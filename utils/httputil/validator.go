package httputil

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"soul/utils"
	"strings"
)

func ParseValidateError(err error, obj any) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors.New("数据格式有误")
	}

	var errResult []string
	for _, e := range errs {
		errStr, err := utils.GetTagValue(obj, e.Field(), e.Tag()+"_err")
		if err != nil {
			// 没有获取到就获取默认的msg
			errStr, _ = utils.GetTagValue(obj, e.Field(), "msg")
		}
		errResult = append(errResult, errStr)

	}
	return errors.New(strings.Join(errResult, ","))
}
