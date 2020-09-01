package conf

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func Test() {
	config := struct {
		Mysql
	}{}
	err := gcfg.ReadFileInto(&config, "utils/conf/ini/conf.ini")
	if err != nil {
		fmt.Println("read error is ", err)
	}
	fmt.Println(config.Mysql)
}
