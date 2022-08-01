/**
 * @Author: Cc
 * @Description: SQLServer Client
 * @File: sql_server
 * @Version: 1.0.0
 * @Date: 2022/8/1 17:54
 * @Software : GoLand
 */

package sql_server

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

func InitStart(address, port, username, password, database string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open(fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v", username, password, address, port, database)), config)
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
