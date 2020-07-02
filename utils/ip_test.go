package utils

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"testing"

	"fmt"
)

func TestResolveIPV4address(t *testing.T) {
	country, area, city, Isp, err := ResolveIPV4address("60.205.176.110")
	//country, area, city, Isp, err := ResolveIPV4address("45.130.146.4")
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Info(country)
	logs.Info(area)
	logs.Info(city)
	logs.Info(Isp)
}

// 外网IP地址
func TestGetExternal(t *testing.T) {
	fmt.Println(GetExternal())
}


func TestGetLocalIPAddress(t *testing.T) {
	g := GetLocalIPAddress()
	logs.Info(g)
}
