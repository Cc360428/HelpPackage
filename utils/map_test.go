package utils

import (
	"log"
	"testing"

	"github.com/goinggo/mapstructure"
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
		log.Println("转换错误 可能传入的不是 map", err.Error())
	}
	log.Println(r.Code)
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
		log.Println(err.Error())
	}
	ma, err := JsonToMap(toJson)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(ma)
}
