/**
 * @Author: Cc
 * @Description: 例子参考
 * @File: example
 * @Version: 1.0.0
 * @Date: 2022/7/29 09:45
 * @Software : GoLand
 */

package redisc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ExampleClient *redis.Client
	ctx           = context.Background()
)

func DeleteRedis(key string) error {
	return ExampleClient.Del(ctx, key).Err()
}

// SetRedisStringCode redis 保存验证码 时间为 60秒
func SetRedisStringCode(key string, code interface{}, expiration time.Duration) (err error) {
	err = ExampleClient.Set(ctx, key, code, 0).Err()
	ExampleClient.Expire(ctx, key, expiration)
	if err != nil {
		fmt.Println("redis 插入失败！", err.Error())
	}
	return
}

// SetRedisString redis 保存验证码 时间为 60秒
func SetRedisString(key string, code interface{}) (err error) {
	err = ExampleClient.Set(ctx, key, code, 0).Err()
	if err != nil {
		fmt.Println("redis 插入失败！", err.Error())
	}
	return
}

// GetStringByKey 通过key获取value （type String）
func GetStringByKey(key string) (value string, err error) {
	value, err = ExampleClient.Get(ctx, key).Result()
	return
}

// SetRedisHash 设置值hash
func SetRedisHash(hashKey string, key string, value interface{}) (err error) {
	err = ExampleClient.HSet(ctx, hashKey, key, value).Err()
	//ExampleClient.Expire(hashKey, 3600*time.Second)
	return err
}

// DeleteHashKey 删除 hash key  删除其中的key
func DeleteHashKey(hashKey string, keys ...string) error {
	return ExampleClient.HDel(ctx, hashKey, keys...).Err()
}

// GetHashString hash key in key get value (string)
func GetHashString(hashKey string, key string) (ret string, err error) {
	//rr, err := ExampleClient.Do("HGet", sessionKey, key)
	ret, err = ExampleClient.HGet(ctx, hashKey, key).Result()
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return ret, nil
}

func GetHashInt64(hashKey string, key string) (ret int64, err error) {
	retStr, err := GetHashString(hashKey, key)
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	retInt, err := strconv.Atoi(retStr)
	if err != nil {
		log.Fatal(err.Error())
		return 0, err
	}
	ret = int64(retInt)
	return ret, nil
}

//////////////////////////////////////////////////////////////////
// list 队列模式 单一的队列模式

type ExampleUser struct {
	Id   int64
	Name string
}

var (
	ExampleLPushKeyList = "Cc:LPush:example"
)

// ProducerMessageList ...
// @Description: 生产者
// LPush Cc:LPush:example '{"Id":1659078721254276000,"Name":"Cc"}'
func ProducerMessageList(id int64) {

	//log.Println("开启生产者")

	var pushCom = ExampleUser{Id: id, Name: "Cc"}
	marshal, err := json.Marshal(pushCom)
	if err != nil {
		log.Fatalln("json Marshal Error:", err.Error())
	}

	result := ExampleClient.LPush(ctx, ExampleLPushKeyList, marshal)
	if result.Err() != nil {
		log.Println("LPush Error:", err.Error())
	}

	//log.Println("LPush ok", result)
}

func ConsumerMessageList() {
	for {
		result, err := ExampleClient.BRPop(ctx, 1*time.Second, ExampleLPushKeyList).Result()
		if err != nil {
			log.Println("BRPop Error:", err.Error())
			continue
		}
		if len(result[1]) == 0 {
			log.Println("消息返回错误")
			continue
		}

		var value ExampleUser
		err = json.Unmarshal([]byte(result[1]), &value)
		if err != nil {
			log.Println("json Unmarshal error:", err.Error())
			continue
		}
		log.Println("接受消息", value)
	}
}
func ConsumerMessageProList() {
	for {
		result, err := ExampleClient.BRPop(ctx, 1*time.Second, ExampleLPushKeyList).Result()
		if err != nil {
			log.Println("BRPop Error:", err.Error())
			continue
		}
		if len(result[1]) == 0 {
			log.Println("消息返回错误")
			continue
		}

		var value ExampleUser
		err = json.Unmarshal([]byte(result[1]), &value)
		if err != nil {
			log.Println("json Unmarshal error:", err.Error())
			continue
		}
		log.Println("接受消息Pro", value)
	}
}

//////////////////////////////////////////////////////////////////
// list 发布者订阅

var ExamplePublishKey = "Cc:publish:example"

// ProducerPubSub 发布者（生产者）
// PUBLISH Cc:publish:example '{"Id":1659078721254276000,"Name":"Cc"}'
func ProducerPubSub(id int64) {

	var pushCom = ExampleUser{Id: id, Name: "Cc"}
	marshal, err := json.Marshal(pushCom)
	if err != nil {
		log.Fatalln("json Marshal Error:", err.Error())
	}

	ExampleClient.Publish(ctx, ExamplePublishKey, marshal)
}

// ConsumerPubSub 订阅者（消费者）
func ConsumerPubSub() {
	subscribe := ExampleClient.Subscribe(ctx, ExamplePublishKey)

	//subscribe.ReceiveMessage(ctx)
	ch := subscribe.Channel()
	// 处理消息
	for msg := range ch {
		log.Printf("消费到数据，channel:%s；message:%s\n", msg.Channel, msg.Payload)
	}
	//	for msg := range ch {
	//msg.
	//		var value ExampleUser
	//		err = json.Unmarshal([]byte(msg), &value)
	//		if err != nil {
	//			log.Println("json Unmarshal error:", err.Error())
	//			continue
	//		}
	//	}

}
func ConsumerPubSubPro() {
	subscribe := ExampleClient.Subscribe(ctx, ExamplePublishKey)

	//subscribe.ReceiveMessage(ctx)
	ch := subscribe.Channel()
	// 处理消息
	for msg := range ch {
		log.Printf("Pro 消费到数据，channel:%s；message:%s\n", msg.Channel, msg.Payload)
	}
	//	for msg := range ch {
	//msg.
	//		var value ExampleUser
	//		err = json.Unmarshal([]byte(msg), &value)
	//		if err != nil {
	//			log.Println("json Unmarshal error:", err.Error())
	//			continue
	//		}
	//	}

}
