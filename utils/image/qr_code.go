// 二维码
package image

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"github.com/skip2/go-qrcode"
	"image/color"
	"image/png"
	"os"
)

// 生成二维码 返回 []byte
func GenerateQrCode(content string) ([]byte, error) {
	encode, err := qrcode.Encode(content, qrcode.Highest, 256)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	return encode, nil
}

// 储存图片
func NewQRCode() {
	code, err := qrcode.New("https://lichaocheng.top", qrcode.Medium)
	if err != nil {
		logs.Error(err.Error())
		return
	}
	// 是否禁用二维码边框
	code.DisableBorder = true
	// 前景色中等颜色。
	//code.ForegroundColor = color.RGBA{R: 0x33, G: 0x33, B: 0x66, A: 0xff}
	// 背景色中等颜色。
	code.BackgroundColor = color.RGBA{R: 0xef, G: 0xef, B: 0xef, A: 0xff}
	// 保存图片 获取image.image
	imagePng := code.Image(256)
	file, err := os.Create("new_qr_code.png")
	if err != nil {
		logs.Error(err.Error())
	}
	defer file.Close()
	err = png.Encode(file, imagePng)
}
