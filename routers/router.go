package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/customer", &controllers.CustomerController{},"delete:DeleteById;get:GetById")
	beego.Router("/customer/name", &controllers.CustomerController{},"delete:DeleteByName")
    beego.Include(&controllers.CustomerController{})
}
