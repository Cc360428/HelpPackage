// Cc Help
package main

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"time"
)

func Cc() {
	logs.Info("Hello 世界")
	logs.Info("Version: v0.0.9")
}

func main() {
	Cc()
	time.Sleep(time.Second * 3)
}
