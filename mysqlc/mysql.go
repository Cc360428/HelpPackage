/**
 * @Author: Cc
 * @Description: Mysql Client
 * @File: mysql
 * @Version: 1.0.0
 * @Date: 2021/8/1 17:10
 * @Software : GoLand
 */

package mysqlc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitStart(address, port, username, password, database string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, address, port, database)), config)

	if err != nil {
		log.Fatalln("open Error", err.Error())
	}

	if s, errDB := db.DB(); errDB != nil {
		log.Fatalln("db Error", err.Error())
	} else {
		if errPing := s.Ping(); errPing != nil {
			log.Fatalln("ping Error", errPing)
		}
	}

	return db, err

}
