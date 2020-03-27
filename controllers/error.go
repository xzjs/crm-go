package controllers

import (
	"bytes"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// ErrorController operations for Error
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error400() {
	c.Ctx.Output.SetStatus(400)
	c.ServeJSON()
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Ctx.Request.Body)
	body := buf.String()
	body = strings.Replace(body, " ", "", -1)
	body = strings.Replace(body, "\n", "", -1)
	logs.Notice(c.Ctx.Request.URL,
		c.Ctx.Request.Method,
		c.Ctx.Request.Form,
		body,
		c.Data["json"])
}

func (c *ErrorController) Error500() {
	c.Ctx.Output.SetStatus(500)
	c.ServeJSON()
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Ctx.Request.Body)
	body := buf.String()
	body = strings.Replace(body, " ", "", -1)
	body = strings.Replace(body, "\n", "", -1)
	logs.Error(c.Ctx.Request.URL,
		c.Ctx.Request.Method,
		c.Ctx.Request.Form,
		body,
		c.Data["json"])
}
