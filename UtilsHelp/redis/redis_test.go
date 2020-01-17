package redis

import (
	"fmt"
	"log"
	"testing"
)

// Test 测试
func Test(t *testing.T) {
	//client, err := NewClient("lichaocheng.top:6379", "lichaocheng", 0)
	client, err := NewClient("127.0.0.1:6379", "", 1)
	if err != nil {
		log.Println("错误：", err.Error())
	}
	fmt.Println("生成的client", client.RedisClient)

}

/**
	fmt.Println("生成的client", client.RedisClient)
	_ = client.SetHash("token", "name", "lcc")
	_ = client.SetHash("token", "age", 23)
	_ = client.SetHash("token", "sex", "m")
	_, _ = client.GetHashString("token", "name")
	name, _ := client.GetHashString("token", "name")
	fmt.Println("name token", name)
	_ = client.SetHash("token", "age", 18)
	age, _ := client.GetHashInt64("token", "age")
	fmt.Println("name token", age)
	time.Sleep(time.Second * 10)
	_ = client.DeleteHashKey("token", "age")
	time.Sleep(time.Second * 10)
	_ = client.Delete("token")
	client.Expire("token", time.Second*30)

	err = client.SetString("name", "lcc")
	if err != nil {
		log.Println("error", err.Error())
	}
	client.Expire("name", time.Second*30)

	err = client.SetString("name", "lcc")
	if err != nil {
		log.Println("error", err.Error())
	}

	value, err := client.GetString("name")
	if err != nil {
		log.Println("error", err.Error())
	}
	fmt.Println("get key value", value)
	//err = client.SetStringTTL("age", 23, time.Second*15)
	//if err != nil {
	//	log.Println("error", err.Error())
	//}

	//time.Sleep(time.Second * 10)
	//err = client.Delete("name")
	//if err != nil {
	//	log.Println("error", err.Error())
	//}
 */
