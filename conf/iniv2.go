package conf

import (
	"fmt"
	"log"

	"github.com/Unknwon/goconfig"
)

func Test1() {
	ini, err := goconfig.LoadConfigFile("utils/conf/ini/conf.ini")
	if err != nil {
		fmt.Println("file error is ", err.Error())
	}
	password, err := ini.GetValue("mysql", "password")
	if err != nil {
		fmt.Println("get error ", err.Error())
	}
	fmt.Println(password)

	redis, err := ini.GetSection("redis")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(redis)
}
