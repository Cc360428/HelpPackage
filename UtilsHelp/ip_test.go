package UtilsHelp

import (
	"testing"

	"fmt"
)

// 
func TestResolveIPV4address(t *testing.T){
	ResolveIPV4address("192.168.1.1")
}

// 外网IP地址
func TestGetExternal(t *testing.T){
	fmt.Println(GetExternal())
}
