package other

import (
	"encoding/json"
	"os"
)

// HelperConvetInterface converts an interface to another interface using JSON marshaling
func HelperConvetInterface(in interface{}, out interface{}) error {
	inBytes, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(inBytes, out)
}

// 取值
func SwaggerJson(swagger string) (v map[string]interface{}) {
	xxx := make(map[string]interface{})
	bytes, err := os.ReadFile(swagger)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(bytes, &xxx)
	if err != nil {
		return nil
	}
	birds := xxx["paths"].(map[string]interface{})
	return birds
}

// 获取值
func GetOpId(vv interface{}) (ret interface{}) {
	dd := make(map[string]interface{})
	_ = HelperConvetInterface(vv, &dd)
	for kk, value := range dd {
		if kk == "operationId" {
			ret = value
			return value
		} else {
			ret = GetOpId(value)
			if ret != nil {
				return ret
			}
		}
	}
	return nil
}
