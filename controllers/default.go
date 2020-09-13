package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hello/models"
)

type MainController struct {
	beego.Controller
}
type CustomerController struct {
	beego.Controller
}

//同时大家注意到新版本里面增加了 URLMapping 这个函数，这是新增加的函数，用户如果没有进行注册，那么就会通过反射来执行对应的函数，如果注册了就会通过 interface 来进行执行函数，性能上面会提升很多。

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// @router /customer/:id [GET]
func (c *CustomerController)GetById() {
	id,_:=c.GetInt(":id")
	fmt.Println(id)
	cc:=models.GetById(id);
	result,err:=json.Marshal(cc)
	if err!=nil{
		return
	}
	c.Ctx.WriteString(string(result));
}

// @router /cc [GET]
func (c *CustomerController)GetByName() {
	//name:=c.Ctx.Input.Param("name")
	name:=c.GetString("name")
	cc:=models.GetByName(name);
	result,err:=json.Marshal(cc)
	if err!=nil{
		return
	}
	c.Ctx.WriteString(string(result));
}

// @router /customer [POST]
func (c *CustomerController)SaveCustomer() {
	body:=c.Ctx.Input.RequestBody
	cc:=models.Customers{}
	json.Unmarshal(body,&cc)
	fmt.Println(cc)
	models.Insert(&cc)
	c.Ctx.WriteString("save success");
}

func (c *CustomerController)DeleteById() {
	id,_:=c.GetInt(":id")
	fmt.Println(id)
	err:=models.DeleteById(id)
	if err!=nil{
		c.Ctx.WriteString("delete failed");
	}else {
		c.Ctx.WriteString("delete success");
	}
}

func (c *CustomerController)DeleteByName() {
	name:=c.GetString("name")
	fmt.Println(name)
	err:=models.DeleteByName(name)
	if err!=nil{
		c.Ctx.WriteString("delete failed");
	}else {
		c.Ctx.WriteString("delete success");
	}
}

