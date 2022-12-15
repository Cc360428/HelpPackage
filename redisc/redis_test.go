/**
 * @Author: Cc
 * @Description: 演示
 * @File: redis_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/28 18:28
 * @Software : GoLand
 */

package redisc

import (
	"context"
	"github.com/Cc360428/HelpPackage/randc"
	"github.com/go-redis/redis/v8"
	"log"
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

func TestLock(t *testing.T) {
	client, err := StartClient("172.12.10.189:6000", "", 6)
	if err != nil {
		t.Log(err.Error())
	}

	var (
		key    = "lcc"
		number = "number"
		//lockKey = "lock_key"
	)

	hSet := client.HSet(ctx, key, number, 15)
	if hSet.Err() != nil {
		t.Log("hset error", hSet.Err())
	}

	go func() {
		for i := 0; i < 220; i++ {
			SetNumber(client, key, number, "A")
		}
	}()

	go func() {

		for i := 0; i < 220; i++ {
			SetNumber(client, key, number, "B")
		}
	}()

	//ch := make(chan bool)
	//<-ch

	time.Sleep(time.Second * 3)
}

func SetNumber(client *redis.Client, key, number, ch string) {
	uuid := randc.GetStringOrder()
	result, SetEX := client.SetNX(ctx, "lockKey", uuid, time.Second*1).Result()
	if SetEX != nil {
		log.Println("SetNX", SetEX.Error())
	}

	if !result {
		//log.Println("result", result)
		return
	}

	defer func() {
		s, err := client.Get(ctx, "lockKey").Result()
		if err != nil {
			return
		}

		if s == uuid {
			client.Del(ctx, "lockKey")
		}
	}()

	get, errHGet := client.HGet(ctx, key, number).Int64()

	if errHGet != nil {
		log.Println("HGet error", errHGet.Error())
	}
	if get <= 0 {
		//fmt.Println(ch, "没库存了")
		return
	}
	//log.Println(get)

	get--

	by := client.HSet(ctx, key, number, get)
	if by.Err() != nil {
		log.Println("HIncrBy error:", by.Err())
	}

	log.Println(ch, "剩余库存", get)
}

//func SetNumber(client *redis.Client, key, number, ch string, l *sync.Mutex) {
//	l.Lock()
//	defer l.Unlock()
//	get, errHGet := client.HGet(ctx, key, number).Int64()
//
//	if errHGet != nil {
//		log.Println("HGet error", errHGet.Error())
//	}
//	if get <= 0 {
//		//fmt.Println(ch, "没库存了")
//		return
//	}
//	//log.Println(get)
//
//	get--
//
//	by := client.HSet(ctx, key, number, get)
//	if by.Err() != nil {
//		log.Println("HIncrBy error:", by.Err())
//	}
//
//	log.Println(ch, "剩余库存", get)
//}

func DeleteNumber(client *redis.Client, key, number, ch string) {
	by := client.HIncrBy(ctx, key, number, -1)
	if by.Err() != nil {
		log.Println("HIncrBy error:", by.Err())
	}

	result, err := by.Result()
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(ch, "剩余库存", result)
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
