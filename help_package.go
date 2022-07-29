// Cc Help
package main

import (
	"log"
	"time"
)

func Cc() {
	log.Println("Hello 世界")
	log.Println("Version: v0.0.9")
}

func main() {
	Cc()
	time.Sleep(time.Second * 3)
}
