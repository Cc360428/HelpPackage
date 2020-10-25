package redis

import (
	"context"
	"github.com/Cc360428/HelpPackage/utils/logs"
	"github.com/go-redis/redis/v8"
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
	_, err = client.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}
	return client, err
}

//SetString 添加string 类型
func (client *Client) SetString(key string, value interface{}) error {
	err := client.RedisClient.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (client *Client) Delete(key string) error {
	err := client.RedisClient.Del(context.Background(), key).Err()
	return err
}

// 设置TTL
func (client *Client) Expire(key string, expiration time.Duration) {
	client.RedisClient.Expire(context.Background(), key, expiration)
}

// 通过key获取value   getType String
func (client *Client) GetString(key string) (value string, err error) {
	value, err = client.RedisClient.Get(context.Background(), key).Result()
	return
}

// 设置值hash
func (client *Client) SetHash(hashKey, key string, value interface{}) (err error) {
	err = client.RedisClient.HSet(context.Background(), hashKey, key, value).Err()
	return err
}

/**
删除 hash key  删除其中的key
*/
func (client *Client) DeleteHashKey(hashKey string, keys ...string) error {
	return client.RedisClient.HDel(context.Background(), hashKey, keys...).Err()
}

/***
hash key in key get value (string)
*/
func (client *Client) GetHashString(hashKey, key string) (ret string, err error) {
	//rr, err := redisClient.Do("HGet", sessionKey, key)
	ret, err = client.RedisClient.HGet(context.Background(), hashKey, key).Result()
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

/********************************** ZSet ***********************************************************************/
// ZSet
// return 1 成功，错误
func (client *Client) ZSet(key string, members ...*redis.Z) (int64, error) {
	if value, err := client.RedisClient.ZAdd(context.Background(), key, members...).Result(); err != nil {
		return 0, err
	} else {
		return value, nil
	}
}

// 查看 zset 总数
// return 数量，错误
func (client *Client) ZCard(key string) (int64, error) {
	if result, err := client.RedisClient.ZCard(context.Background(), key).Result(); err != nil {
		logs.Error(err.Error())
		return 0, err
	} else {
		return result, nil
	}
}

//  命令用于移除有序集中，指定排名(rank)区间内的所有成员
// return 一共移除多少个数量，error
func (client *Client) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	if result, err := client.RedisClient.ZRemRangeByRank(context.Background(), key, start, stop).Result(); err != nil {
		logs.Error(err.Error())
		return 0, err
	} else {
		return result, nil
	}
}
