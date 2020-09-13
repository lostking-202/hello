package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["hello/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hello/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "GetByName",
            Router: "/cc",
            AllowHTTPMethods: []string{"GET"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hello/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "SaveCustomer",
            Router: "/customer",
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:CustomerController"] = append(beego.GlobalControllerRouter["hello/controllers:CustomerController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: "/customer/:id",
            AllowHTTPMethods: []string{"GET"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
