/**
 * @Author: Cc
 * @Description: 描述
 * @File: kafka_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/30 16:54
 * @Software : GoLand
 */

package kafkac

import (
	"log"
	"testing"
	"time"
)

func TestInitStart(t *testing.T) {
	exampleClient, err := InitStartProducer([]string{"172.12.10.189:9092"})
	if err != nil {
		log.Fatalln(err.Error())
	}

	ExampleProducerClient = exampleClient

	tt := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-tt.C:
			ExampleSendMessage(ExampleMessage{
				Id:   time.Now().Unix(),
				Name: "Cc",
			})
		}
	}

}

func TestInitStartP(t *testing.T) {

	consumer, err := InitStartConsumer([]string{"172.12.10.189:9092"}, "1")

	if err != nil {
		log.Println("Marshal Error", err)
		return
	}
	ExampleConsumerClient = consumer

	go ExampleConsume()
	consumerPro, err := InitStartConsumer([]string{"172.12.10.189:9092"}, "2")

	if err != nil {
		log.Println("Marshal Error", err)
		return
	}
	ExampleConsumerClientPro = consumerPro

	go ExampleConsumePro()

	ch := make(chan bool)
	<-ch
}
