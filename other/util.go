package other

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/Cc360428/HelpPackage/uuid"
	"github.com/goinggo/mapstructure"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//  将任何数值转换为Int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	case float64, float32:
		d = int64(val.Float())
	case string:
		d, _ = strconv.ParseInt(val.String(), 10, 64)
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

// 类型断言转类型
func ToInt64V2(v interface{}) (d int64, err error) {
	if _, ok := v.(int); ok {
		d = int64(v.(int))
	} else if _, ok := v.(string); ok {
		d, _ = strconv.ParseInt(v.(string), 10, 64)
	} else if _, ok := v.(float64); ok {
		d = int64(v.(float64))
	} else {
		err = fmt.Errorf("%d 未找到", v)
	}
	return
}

//获取盐值
func Salt() (salt string, err error) {
	UUID, err := uuid.NewV4()
	return UUID.String(), err
}

//string 转 int
func StringTurnInt(pr string) (r int, err error) {
	i, err := strconv.Atoi(pr)
	if err != nil {
		return 0, err
	}
	return i, err
}

// int 转 string
func IntTurnString(pr int) string {
	str := strconv.Itoa(pr)
	return str
}

// int64 转 string
func Int64TurnString(pr int64) string {
	str := strconv.FormatInt(pr, 10)
	return str
}

//interface 转 string
func InterfaceTurnString(pr interface{}) string {
	s := pr.(string)
	return s
}

// map 转 struct
func MapTurnStruct(m map[string]interface{}) (obj interface{}, err error) {
	err = mapstructure.Decode(m, &obj)
	if err != nil {
		err = fmt.Errorf("转换错误！")
	}
	return obj, err
}

// struct 转 map
func StructuralTurnMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

// interface{} 转化为 map
func ConvertInterfaceToMap(src interface{}) (dest map[string]interface{}, isMap bool) {
	isMap = false
	dest = map[string]interface{}{}
	v := reflect.ValueOf(src)
	if v.Kind() != reflect.Map {
		return
	}
	for _, key := range v.MapKeys() {
		dest[key.String()] = v.MapIndex(key).Interface()
	}
	isMap = true
	return
}

// GetBinAbsPath 获取当前运行路径
func GetBinAbsPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

// 对比两者是否相等
func Equals(a, b interface{}) (bool, error) {
	aJson, err := json.Marshal(a)
	if err != nil {
		return false, err
	}
	bJson, err := json.Marshal(b)
	if err != nil {
		return false, err
	}
	if string(aJson) == string(bJson) {
		return true, nil
	} else {
		return false, nil
	}
}

// int64
func DuplicateRemoval(a []int64) []int64 {
	var response []int64
	tmp := make(map[int64]bool)
	for _, value := range a {
		if _, ok := tmp[value]; !ok {
			response = append(response, value)
			tmp[value] = true
		}
	}
	return response
}

// IntChipToStringAll ...
// @Description: 自定义
// @param chip
// @return string
func IntChipToStringAll(toType language.Tag, chip int64) string {
	return message.NewPrinter(toType).Sprintf("%.2v", chip)
}
