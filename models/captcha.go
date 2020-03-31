package models

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// Captcha 图形验证码
type Captcha struct {
	Id   string
	B64s string
}

var store = base64Captcha.DefaultMemStore

// GetCaptcha 获取验证码
func GetCaptcha() (id string, b64s string, err error) {
	bgcolor := color.RGBA{0, 0, 0, 0}
	fonts := []string{"wqy-microhei.ttc"}
	driver := base64Captcha.NewDriverMath(40, 102, 0, 0, &bgcolor, fonts)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = captcha.Generate()
	return id, b64s, err
}

// VerifyCaptcha 校验图形验证码
func VerifyCaptcha(id string, b64s string) bool {
	return store.Verify(id, b64s, true)
}
