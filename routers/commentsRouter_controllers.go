package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["docket-beego/controllers:OauthController"] = append(beego.GlobalControllerRouter["docket-beego/controllers:OauthController"],
        beego.ControllerComments{
            Method: "CreateToken",
            Router: "/token",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["docket-beego/controllers:TaskController"] = append(beego.GlobalControllerRouter["docket-beego/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetTask",
            Router: "/task",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("apiVersion"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["docket-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["docket-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("apiVersion"),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["docket-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["docket-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("apiVersion"),
				param.New("newUserBody", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
