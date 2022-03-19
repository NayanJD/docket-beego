// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"docket-beego/controllers"
	"docket-beego/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.BConfig.RecoverFunc = utils.RecoverPanicFunc

	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSNamespace("/oauth",
				beego.NSInclude(
					&controllers.OauthController{},
				),
			),
			beego.NSNamespace("/task",
				beego.NSInclude(
					controllers.NewTaskController(),
				),
			),
			beego.NSNamespace("/user", beego.NSInclude(controllers.NewUserController())),
		),
	)
	beego.AddNamespace(ns)
}
