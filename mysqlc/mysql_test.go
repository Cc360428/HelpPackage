/**
 * @Author: Cc
 * @Description: 描述
 * @File: mysql_test.go
 * @Version: 1.0.0
 * @Date: 2022/8/1 17:37
 * @Software : GoLand
 */

package mysqlc

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"testing"
	"time"
)

func TestInitStart(t *testing.T) {

	var (
		address  = "172.12.10.189"
		port     = "3306"
		username = "root"
		password = "lichaocheng"
		database = "test"
	)

	dbConfig := new(gorm.Config)

	dbConfig.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "t_",
		SingularTable: true,
		NameReplacer:  nil,
		NoLowerCase:   false,
	}

	dbConfig.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		})

	initDB, err := InitStart(address, port, username, password, database, dbConfig)
	if err != nil {
		t.Error(err)
	}

	// 自动创建表
	err = initDB.AutoMigrate(&User{}, &User{})
	if err != nil {
		fmt.Println("生成数据表错误", err)
	}

	exampleClient = initDB

	InitUserProfile()

	tt := time.NewTicker(time.Second * 1)
	go func() {

		for {
			select {
			case <-tt.C:
				exampleClient.Omit("Name", "Age", "CreatedAt").Create(&User{
					Name:                   "Cc",
					Sex:                    0,
					Avatar:                 time.Now().String(),
					Birthday:               time.Now().String(),
					Email:                  time.Now().String(),
					Password:               time.Now().String(),
					Salt:                   time.Now().String(),
					CreatedAt:              time.Time{},
					UpdatedAt:              time.Time{},
					DeletedAExampleClientt: gorm.DeletedAt{},
				})
			}
		}
	}()
	ch := make(chan bool)
	<-ch
}
