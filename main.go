package main

import (
	"github.com/astaxie/beego/orm"
	"hello/models"
	_ "hello/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
	orm.Debug = true
}

