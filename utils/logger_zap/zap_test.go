/**
 * @Author: Cc
 * @Description: 日志演示
 * @File: zap_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/29 10:06
 * @Software : GoLand
 */

package logger_zap

import (
	"testing"
)

func TestInitZapLog(t *testing.T) {
	logger, err := InitZapLog()
	if err != nil {
		t.Log(err.Error())
	}
	logger.Debug("sada")
	logger.Info("sada")
	logger.Error("error")
}
