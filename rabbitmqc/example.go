/**
 * @Author: Cc
 * @Description: 消息队列演示
 * @File: example
 * @Version: 1.0.0
 * @Date: 2022/7/30 10:32
 * @Software : GoLand
 */

package rabbitmqc

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

var (
	ExampleClient *amqp.Connection
)

type RabbitMQMessage struct {
	Id   int64
	Name string
}

/////////////////////////////////// 消息队列 ///////////////////////////////////

func Producer(queueName string, body RabbitMQMessage) {
	channel, err := ExampleClient.Channel()

	if err != nil {
		log.Println(err.Error())
		return
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Println("QueueDeclare Error", err.Error())
		return
	}

	marshal, err := json.Marshal(body)
	if err != nil {
		log.Println("json Marshal Error", err.Error())
		return
	}

	err = channel.Publish("", queue.Name, false, false, amqp.Publishing{
		Headers:     nil,
		ContentType: "text/plain",
		//ContentEncoding: "",
		//DeliveryMode:    0,
		//Priority:        0,
		//CorrelationId:   "",
		//ReplyTo:         "",
		//Expiration:      "",
		//MessageId:       "",
		//Timestamp:       time.Time{},
		//Type:            "",
		//UserId:          "",
		//AppId:           "",
		Body: marshal,
	})
	if err != nil {
		log.Println("Publish Error", err.Error())
		return
	}
	//log.Println("publish 成功！")
}

func Consumer(queueName string) {
	channel, err := ExampleClient.Channel()
	if err != nil {
		log.Println("Channel Error", err.Error())
		return
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Println("QueueDeclare Error", err.Error())
		return
	}

	consume, err := channel.Consume(queue.Name, "notPro", true, false, false, false, nil)
	if err != nil {
		log.Println("Consume Error", err.Error())
		return
	}

	go func() {
		for d := range consume {
			var mess RabbitMQMessage
			err2 := json.Unmarshal(d.Body, &mess)
			if err2 != nil {
				log.Println("Unmarshal Error", err.Error())
			}
			log.Println("接收的消息", mess)
		}
	}()
	ch := make(chan bool)
	<-ch
}

func ConsumerPro(queueName string) {
	channel, err := ExampleClient.Channel()
	if err != nil {
		log.Println("Channel Error", err.Error())
		return
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Println("QueueDeclare Error", err.Error())
		return
	}

	consume, err := channel.Consume(queue.Name, "pro", true, false, false, false, nil)
	if err != nil {
		log.Println("Consume Error", err.Error())
		return
	}

	go func() {
		for d := range consume {
			var mess RabbitMQMessage
			err2 := json.Unmarshal(d.Body, &mess)
			if err2 != nil {
				log.Println("Unmarshal Error", err.Error())
			}
			log.Println("接收的消息 Pro", mess)
		}
	}()
	ch := make(chan bool)
	<-ch
}

/////////////////////////////////// 发布订阅 ///////////////////////////////////

const ExampleKey = "cc_publish"

func Publish(exchangeName string, body RabbitMQMessage) {
	channel, err := ExampleClient.Channel()
	if err != nil {
		log.Println("Channel Error", err.Error())
		return
	}
	defer channel.Close()
	err = channel.ExchangeDeclare( // 创建一个交换机
		exchangeName, // 交换机名字，
		"fanout",     // 类型，fanout 发布订阅模式
		true, false, false, false, nil)
	if err != nil {
		log.Println("ExchangeDeclare Error", err.Error())
		return
	}
	marshal, err := json.Marshal(body)
	if err != nil {
		log.Println("json Marshal Error", err.Error())
		return
	}
	err = channel.Publish(exchangeName, "", false, false, amqp.Publishing{
		//Headers:         nil,
		ContentType: "text/plain", // 消息内容类型，这里是普通文本
		//ContentEncoding: "",
		//DeliveryMode:    0,
		//Priority:        0,
		//CorrelationId:   "",
		//ReplyTo:         "",
		//Expiration:      "",
		//MessageId:       "",
		//Timestamp:       time.Time{},
		//Type:            "",
		//UserId:          "",
		//AppId:           "",
		Body: marshal,
	})
	if err != nil {
		log.Println("Publish Error", err.Error())
		return
	}
	//log.Println("send Ok ", marshal)
}

func ConsumerQ(exchangeName string) {
	channel, err := ExampleClient.Channel()
	if err != nil {
		log.Println("Channel Error", err.Error())
		return
	}
	defer channel.Close()

	// 声明交换机
	err = channel.ExchangeDeclare(
		exchangeName, // 交换机名，需要跟消息发送方保持一致
		"fanout",     // 交换机类型
		true,         // 是否持久化
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		log.Println("ExchangeDeclare Error", err.Error())
		return
	}
	// 声明需要操作的队列
	queue, err := channel.QueueDeclare(
		"LCC", // 队列名字，不填则随机生成一个
		false, // 是否持久化队列
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = channel.QueueBind(queue.Name, "", exchangeName, false, nil)
	if err != nil {
		log.Println("QueueBind Error", err.Error())
		return
	}
	consume, err := channel.Consume(queue.Name, "LCCQ", true, false, false, false, nil)
	if err != nil {
		log.Println("Consume Error", err.Error())
		return
	}

	// 循环消费队列中的消息
	for d := range consume {
		log.Printf("接收消息:%s", d.Body)
	}

}

func ConsumerQPro(exchangeName string) {
	channel, err := ExampleClient.Channel()
	if err != nil {
		log.Println("Channel Error", err.Error())
		return
	}
	defer channel.Close()

	// 声明交换机
	err = channel.ExchangeDeclare(
		exchangeName, // 交换机名，需要跟消息发送方保持一致
		"fanout",     // 交换机类型
		true,         // 是否持久化
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		log.Println("ExchangeDeclare Error", err.Error())
		return
	}

	// 声明需要操作的队列
	queue, err := channel.QueueDeclare(
		"LCc", // 队列名字，不填则随机生成一个
		false, // 是否持久化队列
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = channel.QueueBind(queue.Name, "", exchangeName, false, nil)
	if err != nil {
		log.Println("QueueBind Error", err.Error())
		return
	}
	consume, err := channel.Consume(queue.Name, "LCCQ1", true, false, false, false, nil)
	if err != nil {
		log.Println("Consume Error", err.Error())
		return
	}

	// 循环消费队列中的消息
	for d := range consume {
		log.Printf("接收消息Pro:%s", d.Body)
	}

}
