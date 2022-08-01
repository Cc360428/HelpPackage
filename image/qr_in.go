package image

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

var errImages error

func createAvatar() (image.Image, error) {
	var (
		bgImg      image.Image
		offset     image.Point
		avatarFile *os.File
		avatarImg  image.Image
	)
	bgImg, errImages = createQrCode("https://studygolang.com/pkgdoc")
	if errImages != nil {
		fmt.Println("创建二维码失败:", errImages)
		return nil, errors.New("创建二维码失败")
	}

	avatarFile, errImages = os.Open("golang.png")
	if errImages != nil {
		log.Println(errImages.Error())
	}
	avatarImg, errImages = png.Decode(avatarFile)
	if errImages != nil {
		log.Println(errImages.Error())
	}
	avatarImg = Resize(avatarImg, 64, 64)
	b := bgImg.Bounds()

	// 设置为居中
	offset = image.Pt((b.Max.X-avatarImg.Bounds().Max.X)/2, (b.Max.Y-avatarImg.Bounds().Max.Y)/2)

	m := image.NewRGBA(b)

	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)

	draw.Draw(m, avatarImg.Bounds().Add(offset), avatarImg, image.Point{X: 0, Y: 0}, draw.Over)

	return m, errImages
}

func createQrCode(content string) (img image.Image, err error) {
	var qrCode *qrcode.QRCode
	qrCode, err = qrcode.New(content, qrcode.Highest)
	if err != nil {
		return nil, errors.New("创建二维码失败")
	}
	qrCode.DisableBorder = true
	img = qrCode.Image(256)
	return img, nil
}

func Resize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}
