// iris response Subpackage
package iris_utils

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"github.com/kataras/iris/v12"
	"net/http"
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
	_, err := ctx.JSON(responseBody)
	if err != nil {
		logs.Info("response error ", err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
	return
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
	_, err := ctx.JSON(responseBody)
	if err != nil {
		logs.Info("response error ", err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
	return
}
