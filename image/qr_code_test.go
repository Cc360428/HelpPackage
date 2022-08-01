package image

import (
	"testing"
)

func TestGenerateQrCode(t *testing.T) {
	code, err := GenerateQrCode("https://lichaocheng.top")
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(code)
}

func TestNewQRCode(t *testing.T) {
	NewQRCode()
}
