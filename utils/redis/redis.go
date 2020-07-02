package redis

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/redis.v5"
	"os"
	"strconv"
	"time"
)

type Client struct {
	RedisClient *redis.Client
}

func NewClient(redisIp, redisPassword string, db int) (redisClient *Client, err error) {
	client := new(Client)
	if redisIp == "" {
		logs.Error("redis address is null in conf")
		os.Exit(1)
	}
	client.RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisIp,
		Password: redisPassword, // no password set
		DB:       db,            // use default DB
	})
	_, err = client.RedisClient.Ping().Result()
	if err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}
	return client, err
}

//SetString 添加string 类型
func (client *Client) SetString(key string, value interface{}) error {
	err := client.RedisClient.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (client *Client) Delete(key string) error {
	err := client.RedisClient.Del(key).Err()
	return err
}

// 设置TTL
func (client *Client) Expire(key string, expiration time.Duration) {
	client.RedisClient.Expire(key, expiration)
}

// 通过key获取value   getType String
func (client *Client) GetString(key string) (value string, err error) {
	value, err = client.RedisClient.Get(key).Result()
	return
}

// 设置值hash
func (client *Client) SetHash(hashKey, key string, value interface{}) (err error) {
	err = client.RedisClient.HSet(hashKey, key, value).Err()
	return err
}

/**
删除 hash key  删除其中的key
*/
func (client *Client) DeleteHashKey(hashKey string, keys ...string) error {
	return client.RedisClient.HDel(hashKey, keys...).Err()
}

/***
hash key in key get value (string)
*/
func (client *Client) GetHashString(hashKey, key string) (ret string, err error) {
	//rr, err := redisClient.Do("HGet", sessionKey, key)
	ret, err = client.RedisClient.HGet(hashKey, key).Result()
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	return ret, nil
}

/**
hash key in key get value (int64)
 */
func (client *Client) GetHashInt64(hashKey string, key string) (ret int64, err error) {
	retStr, err := client.GetHashString(hashKey, key)
	if err != nil {
		logs.Error(err.Error())
		return 0, err
	}
	retInt, err := strconv.Atoi(retStr)
	if err != nil {
		logs.Error(err.Error())
		return 0, err
	}
	ret = int64(retInt)
	return ret, nil
}
