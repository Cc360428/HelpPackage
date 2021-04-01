// Cc 帮助
package main

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"time"
)

// Cc
func Cc() {
	logs.Info("Hello 世界")
	logs.Info("Version: v0.0.8")
}

func main() {
	Cc()
	time.Sleep(time.Second * 3)
}
