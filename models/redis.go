package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var Redis cache.Cache

func init() {
	var err error
	fmt.Println(beego.AppConfig.String("redis"))
	Redis, err = cache.NewCache("redis", beego.AppConfig.String("redis"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
