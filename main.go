package main

import (
	"crm-go/controllers"
	_ "crm-go/routers"
	"fmt"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego/logs"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"), 30)

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}

	var FilterUser = func(ctx *context.Context) {
		uid, ok := ctx.Input.Session("uid").(int64)
		fmt.Println(uid, ok)
		if !ok {
			var whiteMap map[string]int
			whiteMap = make(map[string]int)
			whiteMap["/v1/login"] = 1
			whiteMap["/v1/captcha"] = 1
			whiteMap["/v1/sms"] = 1
			if _, ok = whiteMap[ctx.Request.RequestURI]; !ok {
				ctx.ResponseWriter.WriteHeader(401)
				ctx.WriteString("未登录")
				return
			}
		}
	}
	beego.InsertFilter("/v1/*", beego.BeforeRouter, FilterUser)

	beego.ErrorController(&controllers.ErrorController{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/crm.log"}`)
	beego.Run()
}
