package logrus

import (
	"os"
	"testing"
	"time"
)

// Test 测试
func Test(t *testing.T) {
	log := NewLog(1)
	log.Out = os.Stdout
	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
	 log.Out = file
	} else {
	 log.Info("Failed to log to file, using default stderr")
	}
	log.Info("这是")
	log.WithTime(time.Now()).Info("debug")
	log.Error("error")
}
