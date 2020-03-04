// Gin response 响应参数
package GinUtils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
type CodeType int

const (
	// 成功
	RetSuccess CodeType = 0

	// 路径错误
	RetError CodeType = 1

	// 部分错误
	RetWarning CodeType = -1
)

// ResponseError 定义响应错误
type ResponseError struct {
	Code CodeType    `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseErrorBody(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, ResponseError{
		Code: RetError,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccessBody(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseError{
		Code: RetSuccess,
		Msg:  "",
		Data: data,
	})
}

