package utils

import (
	"encoding/json"
	"fmt"
)

// MapToJson map 转 Json
// Json 通过第三方工具 (https://www.sojson.com/json/json2go.html)
// {"code":200,"data":{"age":19,"user":"lcc"},"massge":"ok"}
// TestMapToJson
func MapToJson(m map[string]interface{}) (string, error) {
	mapToJson, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	jsonToString := string(mapToJson)
	return string(jsonToString), nil
}

// JsonToMap
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
	return m, nil
}

// 取出最小的value
func Min(m map[string]int64) (v string) {
	for k1 := range m {
		for k2 := range m {
			if m[k1] < m[k2] {
				v = k1
			}
		}
	}
	return
}

// 去除最大的
func Max(m map[string]int64) (key string) {
	var maxNumber int64
	for k := range m {
		if m[k] > maxNumber {
			maxNumber = m[k]
			key = k
		}
	}
	return
}
