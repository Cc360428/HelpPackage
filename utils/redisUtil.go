package utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/redis.v5"
	"os"
	"strconv"
	"time"
)

var (
	redisClient *redis.Client
)

func RedisInit(redisIp, redisPassword string, db int) {
	var err error
	if redisIp == "" {
		logs.Error("redis address is null in conf")
		os.Exit(1)
	}
	if redisPassword == "" {
		logs.Error("redis password is null in conf")
		os.Exit(1)
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisIp,
		Password: redisPassword, // no password set
		DB:       db,            // use default DB
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}
}

func DeleteRedis(key string) error {
	return redisClient.Del(key).Err()
}

// redis 保存验证码 时间为 60秒
func SetRedisStringCode(key string, code interface{}, expiration time.Duration) (err error) {
	err = redisClient.Set(key, code, 0).Err()
	redisClient.Expire(key, expiration)
	if err != nil {
		fmt.Println("redis 插入失败！", err.Error())
	}
	return
}

// redis 保存验证码 时间为 60秒
func SetRedisString(key string, code interface{}) (err error) {
	err = redisClient.Set(key, code, 0).Err()
	if err != nil {
		fmt.Println("redis 插入失败！", err.Error())
	}
	return
}

// 通过key获取value （type String）
func GetStringByKey(key string) (value string, err error) {
	value, err = redisClient.Get(key).Result()
	return
}

// 设置值hash
func SetRedisHash(hashKey string, key string, value interface{}) (err error) {
	err = redisClient.HSet(hashKey, key, value).Err()
	//redisClient.Expire(hashKey, 3600*time.Second)
	return err
}

/**
删除 hash key  删除其中的key
*/
func DeleteHashKey(hashKey string, keys ...string) error {
	return redisClient.HDel(hashKey, keys...).Err()
}

/***
hash key in key get value (string)
*/
func GetHashString(hashKey string, key string) (ret string, err error) {
	//rr, err := redisClient.Do("HGet", sessionKey, key)
	ret, err = redisClient.HGet(hashKey, key).Result()
	if err != nil {
		logs.Error(err.Error())
		return "", err
	}
	return ret, nil
}

func GetHashInt64(hashKey string, key string) (ret int64, err error) {
	retStr, err := GetHashString(hashKey, key)
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
