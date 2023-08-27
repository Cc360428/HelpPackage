// Package iris_utils iris response Subpackage
package iris_utils

import (
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
)

// https://www.cnblogs.com/xflonga/p/9368993.html

// 返回参数
type ResponseBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// code code
// msg 消息
// data 消息体
func ResponseStatusOK(ctx iris.Context, code int, msg string, data interface{}) {
	ctx.StatusCode(http.StatusOK)
	ctx.ContentType("application/json")
	var responseBody ResponseBody
	responseBody.Code = code
	responseBody.Msg = msg
	responseBody.Data = data
	err := ctx.JSON(responseBody)
	if err != nil {
		log.Println("response error ", err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
}

// code code
// msg 消息
// data 消息体
func ResponseStatusBadRequest(ctx iris.Context, code int, msg string, data interface{}) {
	ctx.StatusCode(http.StatusBadRequest)
	ctx.ContentType("application/json")
	var responseBody ResponseBody
	responseBody.Code = code
	responseBody.Msg = msg
	responseBody.Data = data
	err := ctx.JSON(responseBody)
	if err != nil {
		log.Println("response error ", err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
}
