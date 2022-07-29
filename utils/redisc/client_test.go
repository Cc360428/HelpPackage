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
	"time"
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

func TestChanList(t *testing.T) {

	exampleClient, err := StartClient("172.12.10.189:6000", "", 6)
	if err != nil {
		t.Log(err.Error())
	}
	ExampleClient = exampleClient

	go ConsumerMessageList()    // 消费者
	go ConsumerMessageProList() // 消费者

	go func() { // 生产者
		for i := 0; i < 5; i++ {
			time.AfterFunc(time.Second*time.Duration(i), func() {
				ProducerMessageList(time.Now().UnixNano())
			})
		}
	}()

	ch := make(chan int)
	<-ch
}

func TestPubSubMessage(t *testing.T) {
	exampleClient, err := StartClient("172.12.10.189:6000", "", 6)
	if err != nil {
		t.Log(err.Error())
	}
	ExampleClient = exampleClient
	go func() {
		for i := 0; i < 15; i++ {
			time.AfterFunc(time.Second*time.Duration(i), func() {
				ProducerPubSub(time.Now().Unix())
			})
		}
	}()
	go ConsumerPubSub()
	go ConsumerPubSubPro()
	ch := make(chan int)
	<-ch
}
