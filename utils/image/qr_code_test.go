package image

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"testing"
)

func TestGenerateQrCode(t *testing.T) {
	code, err := GenerateQrCode("https://lichaocheng.top")
	if err != nil {
		logs.Info(err.Error())
		return
	}
	logs.Info(code)
}

func TestNewQRCode(t *testing.T) {
	NewQRCode()
}
