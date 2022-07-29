/**
 * @Author: Cc
 * @Description: mongodb test
 * @File: mongodb_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/29 16:27
 * @Software : GoLand
 */

package mongodbc

import (
	"testing"
)

func TestInitStart(t *testing.T) {
	//var ctx = context.Background()
	client, err := InitStart("mongodb://root:root@172.12.10.189:27001")
	if err != nil {
		t.Log(err)
	}

	ExampleClient = client

	InsertOne("cc", "cc")
	DeleteCollection("cc", "lcc")
	// 删除
	DeleteDataBase("cc")
}
