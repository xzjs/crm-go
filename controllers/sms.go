package controllers

import (
	"crm-go/models"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
)

type SmsController struct {
	beego.Controller
}

// @Title Post
// @Description 发送验证码
// @Param body body models.Sms true "手机号和验证码"
// @Success 200 {string} success
// @Failure 400 {string} 图形验证码错误
// @router / [post]
func (c *SmsController) Post() {
	var sms models.Sms
	json.Unmarshal(c.Ctx.Input.RequestBody, &sms)
	if !models.VerifyCaptcha(sms.Captcha.Id, sms.Captcha.B64s) {
		c.Data["json"] = "图形验证码错误"
		c.Abort("400")
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	vcode = "111111" // 测试代码
	models.Redis.Put(sms.Mobile, vcode, 10*time.Minute)

	// TODO 短信发送

	c.Data["json"] = "success"
	c.ServeJSON()
}
