package controllers

import (
	"crm-go/models"
	"encoding/json"
	"image/color"

	"github.com/astaxie/beego"
	"github.com/mojocn/base64Captcha"
)

type CaptchaController struct {
	beego.Controller
}

var store = base64Captcha.DefaultMemStore

// @Title GetCaptcha
// @Description 获取图形验证码
// @Success 200 {object} models.Captcha
// @router / [get]
func (c *CaptchaController) Get() {
	bgcolor := color.RGBA{0, 0, 0, 0}
	fonts := []string{"wqy-microhei.ttc"}
	driver := base64Captcha.NewDriverMath(40, 102, 0, 0, &bgcolor, fonts)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = models.Captcha{id, b64s}
	}
	c.ServeJSON()
}

// @Title VerifyCaptcha
// @Description 校验图形验证码
// @Param body body models.Captcha true "验证码"
// @Success 200 {bool} true or false
// @router / [post]
func (c *CaptchaController) Post() {
	var captcha models.Captcha
	json.Unmarshal(c.Ctx.Input.RequestBody, &captcha)
	if store.Verify(captcha.ID, captcha.B64s, true) {
		c.Data["json"] = true
	} else {
		c.Data["json"] = false
	}
	c.ServeJSON()
}
