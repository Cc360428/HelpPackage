/**
 * @Author: Cc
 * @Description: redis 客户端
 * @File: client
 * @Version: 1.0.0
 * @Date: 2022/7/28 18:23
 * @Software : GoLand
 */

package redisc

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// StartClient ...
// @Description: 初始化 ExampleClient 给出使用者最大权利
// @param redisIp
// @param redisPassword
// @param db
// @return *redis.Client
// @return error
func StartClient(redisIp, redisPassword string, db int) (*redis.Client, error) {

	if redisIp == "" {
		log.Println("redis address is null in conf")
		return nil, errors.New("redisIp is null")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisIp,
		Password: redisPassword, // no password set
		DB:       db,            // use default DB
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("redis ping error:%v", err.Error())
	}

	return redisClient, nil
}
