/**
 * @Author: Cc
 * @Description: mongodb
 * @File: mongodb
 * @Version: 1.0.0
 * @Date: 2022/7/29 16:17
 * @Software : GoLand
 */

package mongodbc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
)

func InitStart(url string) (*mongo.Client, error) {

	var ctx = context.Background()
	connect, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalln("applyUrl Error:", err.Error())
	}

	err = connect.Ping(ctx, nil)
	if err != nil {
		log.Fatalln("Ping Error:", err)
		return nil, err
	}

	return connect, nil

}
