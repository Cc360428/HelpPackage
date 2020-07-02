package utils

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"testing"
)

func TestGetStringOrder(t *testing.T) {
	order := GetStringOrder()
	logs.Info(order)

}
