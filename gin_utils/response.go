// Gin response 响应参数
package gin_utils

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

// ResponseErrorBody 请求错误
func ResponseErrorBody(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, ResponseError{
		Code: RetError,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccessBody 请求成功
func ResponseSuccessBody(c *gin.Context, msg, data interface{}) {
	c.JSON(http.StatusOK, ResponseError{
		Code: RetSuccess,
		Msg:  msg,
		Data: data,
	})
}

// ResponseWarningBody 警告
func ResponseWarningBody(c *gin.Context, msg, data interface{}) {
	c.JSON(http.StatusOK, ResponseError{
		Code: RetWarning,
		Msg:  msg,
		Data: data,
	})
}
