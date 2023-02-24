/**
 * @Author: Cc
 * @Description: telegram message
 * @File: telegram_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/27 16:07
 * @Software : GoLand
 */

package monitor_message

import (
	"fmt"
	"testing"
	"time"
)

func TestSendTelegram(t *testing.T) {
	SendTelegram(-1001734918155, "5447964690:AAFHOTD3FaGgVLiOrhgbDRUznQ0bCOsHcY4",
		fmt.Sprintf("%v %v \n", time.Now().Format("2006-01-02 15:04:05"), "HelpPackage Test Send Message Telegram"))
}
