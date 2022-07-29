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
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	exampleClient *redis.Client
	ctx           = context.Background()
)

func DeleteRedis(key string) error {
	return exampleClient.Del(ctx, key).Err()
}

// SetRedisStringCode redis 保存验证码 时间为 60秒
func SetRedisStringCode(key string, code interface{}, expiration time.Duration) (err error) {
	err = exampleClient.Set(ctx, key, code, 0).Err()
	exampleClient.Expire(ctx, key, expiration)
	if err != nil {
		fmt.Println("redis 插入失败！", err.Error())
	}
	return
}

// SetRedisString redis 保存验证码 时间为 60秒
func SetRedisString(key string, code interface{}) (err error) {
	err = exampleClient.Set(ctx, key, code, 0).Err()
	if err != nil {
		fmt.Println("redis 插入失败！", err.Error())
	}
	return
}

// GetStringByKey 通过key获取value （type String）
func GetStringByKey(key string) (value string, err error) {
	value, err = exampleClient.Get(ctx, key).Result()
	return
}

// SetRedisHash 设置值hash
func SetRedisHash(hashKey string, key string, value interface{}) (err error) {
	err = exampleClient.HSet(ctx, hashKey, key, value).Err()
	//exampleClient.Expire(hashKey, 3600*time.Second)
	return err
}

// DeleteHashKey 删除 hash key  删除其中的key
func DeleteHashKey(hashKey string, keys ...string) error {
	return exampleClient.HDel(ctx, hashKey, keys...).Err()
}

// GetHashString hash key in key get value (string)
func GetHashString(hashKey string, key string) (ret string, err error) {
	//rr, err := exampleClient.Do("HGet", sessionKey, key)
	ret, err = exampleClient.HGet(ctx, hashKey, key).Result()
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
