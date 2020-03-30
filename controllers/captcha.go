package controllers

import (
	"crm-go/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type CaptchaController struct {
	beego.Controller
}

// @Title GetCaptcha
// @Description 获取图形验证码
// @Success 200 {object} models.Captcha
// @router / [get]
func (c *CaptchaController) Get() {
	id, b64s, err := models.GetCaptcha()
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
	if models.VerifyCaptcha(captcha.ID, captcha.B64s) {
		c.Data["json"] = true
	} else {
		c.Data["json"] = false
	}
	c.ServeJSON()
}
