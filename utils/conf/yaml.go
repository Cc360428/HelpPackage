package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Test3() {
	fmt.Println("yaml 读取文件")
	var c conf
	conf := c.getConf()
	fmt.Println(conf)
}

//profile variables
type conf struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"password"`
	DBName string `yaml:"dbname"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("utils/conf/yaml/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
