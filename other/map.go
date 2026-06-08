package other

import (
	"encoding/json"
	"fmt"
	"sort"
)

// MapToJson map 转 Json
// Json 通过第三方工具 (https://www.sojson.com/json/json2go.html)
// {"code":200,"data":{"age":19,"user":"lcc"},"massage":"ok"}
// TestMapToJson
func MapToJson(m map[string]interface{}) (string, error) {
	mapToJson, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	jsonToString := string(mapToJson)
	return string(jsonToString), nil
}

// JsonToMap ...
// @Description:
// @param jsonStr
// @return map[string]interface{}
// @return error
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

// Min ...
// @Description: 取出最小的value
// @param m
// @return key, 是否找到
func Min(m map[string]int64) (key string, found bool) {
	if len(m) == 0 {
		return "", false
	}
	var min int64
	var first bool = true
	for k, v := range m {
		if first {
			min = v
			key = k
			first = false
			continue
		}
		if v < min {
			min = v
			key = k
		}
	}
	return key, true
}

// Max ...
// @Description: 取出最大的
// @param m
// @return key, 是否找到
func Max(m map[string]int64) (key string, found bool) {
	if len(m) == 0 {
		return "", false
	}
	var max int64
	var first bool = true
	for k, v := range m {
		if first {
			max = v
			key = k
			first = false
			continue
		}
		if v > max {
			max = v
			key = k
		}
	}
	return key, true
}

func SortMap(rate map[int64]float64) []int64 {
	keys := make([]int64, 0, len(rate))
	for k := range rate {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return rate[keys[i]] > rate[keys[j]] })
	return keys
}
