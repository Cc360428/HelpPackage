package utils

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"github.com/goinggo/mapstructure"
	"testing"
)

type HttpResponse struct {
	Code int `json:"code"`
	Data struct {
		Age  int    `json:"age"`
		User string `json:"user"`
	} `json:"data"`
	Message string `json:"massge"`
}

func TestMapToJson(t *testing.T) {
	m := make(map[string]interface{})
	m["code"] = 200
	m["massge"] = "ok"
	data := make(map[string]interface{})
	data["user"] = "lcc"
	data["age"] = 19
	m["data"] = data
	var r HttpResponse
	err := mapstructure.Decode(m, &r)
	if err != nil {
		logs.Error("转换错误 可能传入的不是 map", err.Error())
	}
	logs.Info(r.Code)
}

func TestJsonToMap(t *testing.T) {
	m := make(map[string]interface{})
	m["code"] = 200
	m["massge"] = "ok"
	data := make(map[string]interface{})
	data["user"] = "lcc"
	data["age"] = 19
	m["data"] = data
	toJson, err := MapToJson(m)
	if err != nil {
		logs.Error(err.Error())
	}
	ma, err := JsonToMap(toJson)
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Info(ma)
}
