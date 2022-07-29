/**
 * @Author: Cc
 * @Description: mongodb 演示用例
 * @File: example
 * @Version: 1.0.0
 * @Date: 2022/7/29 18:05
 * @Software : GoLand
 */

package mongodbc

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	ExampleClient *mongo.Client
	ExampleCtx    = context.Background()
)

func InsertOne(db, collection string) {

	client := ExampleClient.Database(db).Collection(collection)

	res, err := client.InsertOne(ExampleCtx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("返回数据", res.InsertedID)
}

func DeleteCollection(db, collection string) {
	client := ExampleClient.Database(db).Collection(collection)
	client.Drop(ExampleCtx)
}

// DeleteDataBase 删除数据库
func DeleteDataBase(database string) {
	ExampleClient.Database(database).Drop(ExampleCtx)
}
