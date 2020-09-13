package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Customers struct {
	Id int
	Name string
}

func Init()  {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:root123@/world?charset=utf8");
	orm.RegisterModel(new(Customers))
}

func Insert(user *Customers)  {
	o:=orm.NewOrm()
	id,err:=o.Insert(user)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}

func GetById(id int) []*Customers {
	o:=orm.NewOrm()
	var cs []*Customers
	num,err:=o.QueryTable("customers").Filter("id",id).All(&cs,"id","name");
	if err==nil{
		fmt.Println(num)
	}
	return cs;
}

func GetByName(name string) []*Customers {
	o:=orm.NewOrm()
	var cs []*Customers
	num,err:=o.QueryTable("customers").Filter("name",name).All(&cs,"id","name");
	if err==nil{
		fmt.Println(num)
	}
	return cs;
}

func DeleteById(id int)error{
	o:=orm.NewOrm()
	_,err:=o.Delete(&Customers{Id:id});
	return err
}

//基本的增删改查只能操作主键
func DeleteByName(name string)error{
	orm.Debug = true
	o:=orm.NewOrm()
	_,err:=o.Delete(&Customers{Name:name});
	return err
}

//QueryBuilder 原生sql