/**
 * @Author: Cc
 * @Description: SQLServer 演示
 * @File: example
 * @Version: 1.0.0
 * @Date: 2022/8/1 18:01
 * @Software : GoLand
 */

package sql_server

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

var (
	exampleClient *gorm.DB
)

func EInsert() {
	t := time.Now()
	// 插入数据
	sqlStr := fmt.Sprintf("INSERT INTO[dbo].[MLogOnline]([GameId], [GroupId], [Online], [Time], [OS]) VALUES(%v, %v, %v, '%v', 1);", 2020, 1, 666, t.Truncate(time.Second*30).Format("2006-01-02 15:04:05"))
	fmt.Println(sqlStr)
	tx := exampleClient.Exec(sqlStr)
	if tx.Error != nil {
		fmt.Println("Insert data err=", tx.Error)
	}
	fmt.Println("成功插入数据!")
}
