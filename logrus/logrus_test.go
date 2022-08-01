/**
 * @Author: Cc
 * @Description: 日志
 * @File: logrus_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/29 10:16
 * @Software : GoLand
 */

package logrus

import (
	"log"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewLogs(t *testing.T) {
	logs := NewLogs()
	log.Println("test")
	logs.WithFields(logrus.Fields{
		"name":  "Cc",
		"age":   11,
		"email": "li_chao_cheng@163.com",
	}).Info("Failed to sewqnd event")

}
