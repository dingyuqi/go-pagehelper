// Package test -----------------------------
// @file      : respond.go
// @author    : dingyq
// @time      : 2024/7/16 下午5:23
// -------------------------------------------
package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 200
)

func OkWithData(data interface{}, c *gin.Context) {
	result(SUCCESS, data, "OperationSuccess", c)
}

func result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, response{
		code,
		data,
		msg,
	})
}

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
