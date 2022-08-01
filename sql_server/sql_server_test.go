/**
 * @Author: Cc
 * @Description: 描述
 * @File: sql_server_test.go
 * @Version: 1.0.0
 * @Date: 2022/8/1 18:01
 * @Software : GoLand
 */

package sql_server

import (
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestInitStart(t *testing.T) {
	var (
		address  = "172.13.7.24"
		port     = "38000"
		username = "GameUser"
		password = "wDucbhutnU1X"
		database = "OS_Log"
		dbConfig = new(gorm.Config)
	)

	db, err := InitStart(address, port, username, password, database, dbConfig)
	if err != nil {
		log.Fatalln(err)
	}

	exampleClient = db
	EInsert()
}
