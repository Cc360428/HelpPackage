/**
 * @Author: Cc
 * @Description: 消息队列测试
 * @File: rabbitmq_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/30 10:31
 * @Software : GoLand
 */

package rabbitmqc

import (
	"testing"
	"time"
)

// 消息队列，读立马销毁
func Test_initStart(t *testing.T) {
	var url = "amqp://guest:guest@172.12.10.189:5672/"
	client, err := initStart(url)
	if err != nil {
		t.Fatal("initStart Error:", err.Error())
	}

	ExampleClient = client

	go func() {
		tt := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-tt.C:
				Producer("cc", RabbitMQMessage{
					Id:   time.Now().Unix(),
					Name: "Cc",
				})
			}
		}
	}()

	go ConsumerPro("cc")
	go Consumer("cc")

	ch := make(chan bool)
	<-ch
}

// 发布订阅
func Test_initStartPro(t *testing.T) {
	var url = "amqp://guest:guest@172.12.10.189:5672/"
	client, err := initStart(url)
	if err != nil {
		t.Fatal("initStart Error:", err.Error())
	}

	// 发布订阅
	ExampleClient = client

	go func() {
		tt := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-tt.C:
				Publish("ccP", RabbitMQMessage{
					Id:   time.Now().Unix(),
					Name: "Cc",
				})
			}
		}
	}()

	go ConsumerQ("ccP")
	go ConsumerQPro("ccP")

	ch := make(chan bool)
	<-ch
}
