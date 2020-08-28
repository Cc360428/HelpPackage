// beego 参数返回
package beego_utils

import "github.com/astaxie/beego"

type Result struct {
	Code CodeType    `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

type CodeType int

const (
	// 200
	RetSuccess CodeType = 0

	// 404
	RetError CodeType = 1

	// other
	RetWarning CodeType = -1
)

// init
func ResultInit() *Result {
	var result = Result{Code: 1, Data: nil, Msg: ""}
	return &result
}

// 返回错误
func ReturnFail(this *beego.Controller, msg string) {
	var ret Result
	ret.Code = RetError
	ret.Msg = msg
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}

//  返回正确
func ReturnSuccess(this *beego.Controller, data interface{}) {
	var ret Result
	ret.Code = RetSuccess
	ret.Data = data
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}

// 返回告警
func ReturnMuti(this *beego.Controller, msg interface{}, data interface{}) {
	var ret Result
	ret.Code = RetSuccess
	ret.Msg = msg
	ret.Data = data
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}

// 返回告警
func ReturnWarning(this *beego.Controller, msg interface{}, data interface{}) {
	var ret Result
	ret.Code = RetWarning
	ret.Msg = msg
	ret.Data = data
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}
