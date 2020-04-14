// @APIVersion 1.0.0
// @Title CRM API
// @Description CRM API 文档
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"crm-go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/captcha",
			beego.NSInclude(
				&controllers.CaptchaController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/file",
			beego.NSInclude(
				&controllers.FileController{},
			),
		),
		beego.NSNamespace("/sms",
			beego.NSInclude(
				&controllers.SmsController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),
		beego.NSNamespace("/result",
			beego.NSInclude(
				&controllers.ResultController{},
			),
		),
		beego.NSNamespace("/record",
			beego.NSInclude(
				&controllers.RecordController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
