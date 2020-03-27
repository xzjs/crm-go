package main

import (
	"crm-go/controllers"
	_ "crm-go/routers"

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
	beego.ErrorController(&controllers.ErrorController{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/crm.log"}`)
	beego.Run()
}
