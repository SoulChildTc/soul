package httputil

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OK(c *gin.Context, data any, msg string) {
	resp := ResponseBody{
		Status: "success",
		Msg:    msg,
		Data:   data,
	}
	c.JSON(http.StatusBadRequest, resp)
}

func Error(c *gin.Context, msg string) {
	ErrorWithCode(c, http.StatusBadRequest, msg)
}

func ErrorWithCode(c *gin.Context, code int, msg string) {
	resp := ResponseBody{
		Status: "error",
		Msg:    msg,
	}
	c.JSON(code, resp)
}
