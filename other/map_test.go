package other

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

func TestMax(t *testing.T) {
	r := map[string]int64{}
	r["Cc"] = 26
	r["LCC"] = 66
	t.Log(Max(r))
}

func TestMin(t *testing.T) {
	r := map[string]int64{}
	r["Cc"] = 26
	r["LCC"] = 4
	r["XM"] = 1
	r["LZ"] = 2
	r["ZS"] = 3
	t.Log(Min(r))
}
