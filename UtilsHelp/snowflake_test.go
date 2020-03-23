package UtilsHelp

import (
	"github.com/Cc360428/HelpPackage/UtilsHelp/logs"
	"testing"
)

func TestGetStringOrder(t *testing.T) {
	order := GetStringOrder()
	logs.Info(order)

}
