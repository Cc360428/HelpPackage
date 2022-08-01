package image

import (
	"image/png"
	"log"
	"os"
	"testing"
)

func Test_createAvatar(t *testing.T) {
	avatar, err := createAvatar()
	if err != nil {
		log.Println(err.Error())
	}
	file, err := os.Create("qr_in.png")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	err = png.Encode(file, avatar)
	if err != nil {
		t.Log(err.Error())
	}
}
