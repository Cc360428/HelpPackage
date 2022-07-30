/**
 * @Author: Cc
 * @Description: kafka 演示
 * @File: example
 * @Version: 1.0.0
 * @Date: 2022/7/30 17:27
 * @Software : GoLand
 */

package kafkac

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

var (
	ExampleProducerClient    sarama.SyncProducer
	ExampleConsumerClient    sarama.ConsumerGroup
	ExampleConsumerClientPro sarama.ConsumerGroup
	Topic                    = "CcCode"
	ctx                      = context.Background()
)

type ExampleMessage struct {
	Id   int64
	Name string
}

func ExampleSendMessage(message ExampleMessage) {
	if message.Id == 0 {
		return
	}

	marshal, err := json.Marshal(message)
	if err != nil {
		log.Println("Marshal Error", err)
	}
	sendMessage, i, err := ExampleProducerClient.SendMessage(&sarama.ProducerMessage{
		Topic: Topic,
		//Key:       nil,
		Value: sarama.ByteEncoder(marshal),
		//Headers:   nil,
		//Metadata:  nil,
		//Offset:    0,
		//Partition: 0,
		//Timestamp: time.Time{},
	})
	if err != nil {
		log.Println("SendMessage Error", err)
		return
	}
	log.Printf("send OK partition:%v offset:%v\n", sendMessage, i)

}

type consumerGroupHandler struct {
	name string
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		// 手动确认消息
		sess.MarkMessage(msg, "")
	}
	return nil
}

func ExampleConsume() {
	err := ExampleConsumerClient.Consume(ctx, []string{Topic}, consumerGroupHandler{
		name: "Cc",
	})

	if err != nil {

	}
}

type consumerGroupHandlerPro struct {
	name string
}

func (consumerGroupHandlerPro) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandlerPro) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandlerPro) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		// 手动确认消息
		sess.MarkMessage(msg, "")
	}
	return nil
}

func ExampleConsumePro() {
	err := ExampleConsumerClientPro.Consume(ctx, []string{Topic}, consumerGroupHandler{
		name: "CcPro",
	})

	if err != nil {

	}
}
