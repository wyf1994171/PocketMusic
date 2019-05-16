package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回调用者的数据, err_msg里填错误的信息，如果没有问题置为空
type Response struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func writeResponse(c *gin.Context, errCode int, errMsg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    data,
	})
}
