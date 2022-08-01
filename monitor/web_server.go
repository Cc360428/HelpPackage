// 监听文件
package main

import (
	"fmt"
	ginUtils "github.com/Cc360428/HelpPackage/web/gin"
	"log"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	conf        Person
	s           sync.WaitGroup
	MysqlClient *gorm.DB
)

func main() {
	engine := gin.New()
	engine.GET("/", root)
	log.Println("get---> http://localhost:8080")
	s.Add(2)
	go monitor()
	go RConfig()
	s.Wait()
	mysqlClient()
	err := engine.Run(fmt.Sprintf(":%d", conf.HttpPort))
	if err != nil {
		panic(err)
	}
}

type CcUser struct {
	ID            int64     `gorm:"column:id" json:"id" form:"id"`
	Name          string    `gorm:"column:name" json:"name" form:"name"`
	Sex           int64     `gorm:"column:sex" json:"sex" form:"sex"`
	ImagesUrl     string    `gorm:"column:images_url" json:"images_url" form:"images_url"`
	Email         string    `gorm:"column:email" json:"email" form:"email"`
	Password      string    `gorm:"column:password" json:"password" form:"password"`
	Salt          string    `gorm:"column:salt" json:"salt" form:"salt"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	LastLogonTime time.Time `gorm:"column:last_logon_time" json:"last_logon_time" form:"last_logon_time"`
	DeletedAt     time.Time `gorm:"column:deleted_at" json:"deleted_at" form:"deleted_at"`
	UserId        string    `gorm:"column:user_id" json:"user_id" form:"user_id"`
}

func root(context *gin.Context) {
	var users []CcUser
	err := MysqlClient.Raw("select * from cc_user").Scan(&users).Error
	if err != nil {
		log.Println(err.Error())
	}
	ginUtils.ResponseSuccessBody(context, "ok", users)
}

type Person struct {
	HttpPort int    `json:"http_port"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func RConfig() {
	var path = "utils/monitor/conf/base.toml"
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	s.Done()
	log.Println("this -->", conf)
}

func mysqlClient() {
	log.Println("更新mysql")
	mysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", conf.Password,
		"lichaocheng.top",
		conf.Port,
		"cc")
	db, err := gorm.Open("mysql", mysql)
	if err != nil {
		panic(err)
	}
	MysqlClient = db
}
