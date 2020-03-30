package models

import (
	"fmt"

	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var Redis cache.Cache

func init() {
	var err error
	Redis, err = cache.NewCache("redis", `{"key":"crm","conn":"127.0.0.1:6379"}`)
	if err != nil {
		fmt.Println(err.Error())
	}
}
