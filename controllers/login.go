package controllers

import (
	"crm-go/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

// @Title Login
// @Description 登录
// @Param body body models.Login true "登录信息"
// @Success 200 {int} 用户id
// @Failure 400 {string} 验证码错误
// @Failure 500 {string} 服务器错误
// @router / [post]
func (l *LoginController) Post() {
	var login models.Login
	json.Unmarshal(l.Ctx.Input.RequestBody, &login)
	vcode := models.Redis.Get(login.Mobile)
	if login.Vcode == vcode {
		if id, err := models.GetUserByMobile(login.Mobile); err == nil {
			l.SetSession("uid", id)
			l.Data["json"] = id
		} else {
			l.Ctx.ResponseWriter.WriteHeader(500)
			l.Data["json"] = err.Error()
		}
	} else {
		l.Ctx.ResponseWriter.WriteHeader(400)
		l.Data["json"] = "验证码错误"
	}
	l.ServeJSON()
}

// @Title GetLoginUser
// @Description 获取登录用户
// @Success 200 {object} 用户id
// @Failure 401 {string} 验证码错误
// @router / [get]
func (l *LoginController) Get() {
	id := l.GetSession("uid")
	v, err := models.GetUserById(id.(int64))
	if err != nil {
		l.Data["json"] = err.Error()
	} else {
		l.Ctx.ResponseWriter.WriteHeader(401)
		l.Data["json"] = v
	}
	l.ServeJSON()
}
