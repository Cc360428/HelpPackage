package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"log"
	"testing"
)

func Test(t *testing.T) {
	session, err := NewMongdodb("127.0.0.1")
	if err != nil {
		log.Panicln("错误：", err.Error())
	}
	client := session.DB("Test").C("info")

	//创建数据
	data := Student{
		Name:   "lcc",
		Age:    23,
		Sid:    "8888",
		Status: 1,
	}
	err = client.Insert(&data)
	if err != nil {
		log.Println(err.Error(), "插入错误")
	}

	user := new(Student)
	err = client.Find(bson.M{"sid": "8888"}).One(&user)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(user)
}

type Student struct {
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
	Sid    string `bson:"sid"`
	Status int    `bson:"status"`
}

type Per struct {
	Per []Student
}
