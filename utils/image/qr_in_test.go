package image

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"image/png"
	"os"
	"testing"
)

func Test_createAvatar(t *testing.T) {
	avatar, err := createAvatar()
	if err != nil {
		logs.Error(err.Error())
	}
	file, err := os.Create("qr_in.png")
	if err != nil {
		logs.Error(err.Error())
	}
	defer file.Close()
	err = png.Encode(file, avatar)
}
