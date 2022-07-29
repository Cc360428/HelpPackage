package netc

import (
	"log"
	"testing"

	"fmt"
)

func TestResolveIPV4address(t *testing.T) {
	country, area, city, Isp, err := ResolveIPV4address("60.205.176.110")
	//country, area, city, Isp, err := ResolveIPV4address("45.130.146.4")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(country)
	log.Println(area)
	log.Println(city)
	log.Println(Isp)
}

// 外网IP地址
func TestGetExternal(t *testing.T) {
	fmt.Println(GetExternal())
}

func TestGetLocalIPAddress(t *testing.T) {
	g := GetLocalIPAddress()
	log.Println(g)
}
