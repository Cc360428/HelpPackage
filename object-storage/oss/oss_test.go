/**
 * @Author: Cc
 * @Description: 描述
 * @File: oss_test.go
 * @Version: 1.0.0
 * @Date: 2023/7/6 16:14
 * @Software : GoLand
 */

package oss

import (
	"encoding/json"
	"io"
	"testing"
)

func TestNewOSSClient(t *testing.T) {

	client, err := NewOSSClient(Config{
		Endpoint:        "",
		AccessKeyID:     "",
		AccessKeySecret: "",
		BucketName:      "",
		Prefix:          "2020/dev/",
		Suffix:          ".json",
	})
	if err != nil {
		t.Log(err.Error())
	}

	key := "playerId"
	setValue := make([]response, 10)
	for i := 0; i < 10; i++ {
		setValue[i] = response{
			Id:   i,
			Code: "200",
		}
	}

	marshal, err := json.Marshal(setValue)
	if err != nil {
		t.Log(err.Error())
		return
	}

	t.Log(setValue)
	err = client.Put(Object{Key: key, Value: string(marshal)})
	if err != nil {
		t.Log(err.Error())
	}

	//err = client.Delete(key)
	//if err != nil {
	//	t.Log(err.Error())
	//}
	//
	//err = client.ClearObject(key)
	//if err != nil {
	//	t.Log(err.Error())
	//}

	//t.Log(client.key("7197934"))
	//exist, err := client.IsObjectExist(key)
	//if err != nil {
	//	return
	//}
	//getValue, err := client.Get(key)
	//if err != nil {
	//	t.Log(err)
	//}
	//fmt.Println(string(getValue)) // 在控制台输出JSON内容
	//
	//toStruct, err := client.GetToStruct(key)
	//if err != nil {
	//	t.Log(err.Error())
	//	return
	//}
	//
	//dataStruct, err := readToStructs(toStruct)
	//if err != nil {
	//	fmt.Println("Error converting to struct:", err)
	//	return
	//}
	//
	//t.Log(dataStruct)
	//
	//list, err := client.GetObjectList("", 100)
	//if err != nil {
	//	return
	//}
	//
	//t.Log(list)
	expressionSQL := "select * from ossobject [*] s where s.Id > 1 limit 2"
	lines, err := client.GetJsonObjectByLines(key, expressionSQL)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(lines))
	expressionSQL1 := "select person.Id from ossobject"
	lines1, err := client.GetJsonObjectByLines(key, expressionSQL1)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(lines1))
}

func readToStructs(reader io.ReadCloser) (result []*response, err error) {
	defer reader.Close()

	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func readToStruct(reader io.ReadCloser) (*response, error) {
	defer reader.Close()

	var result response
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

type response struct {
	Id   int
	Code string
}
