/**
 * @Author: Cc
 * @Description: 描述
 * @File: client_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/28 18:28
 * @Software : GoLand
 */

package redisc

import (
	"context"
	"testing"
)

func TestStartClient(t *testing.T) {
	client, err := StartClient("172.12.10.189:6000", "", 6)
	if err != nil {
		t.Log(err.Error())
	}
	client.Set(context.Background(), "666", "999", -1)
	get := client.Get(context.Background(), "666")
	t.Log(get)
}
