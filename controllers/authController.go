package controllers

import (
	"docket-beego/auth"
	"docket-beego/models"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
	user          *models.User
	authWhitelist map[string]bool
}

func (ctrl *AuthController) Prepare() {
	tokenInfo, err := auth.Srv.ValidationBearerToken(ctrl.Ctx.Request)

	_, action := ctrl.GetControllerAndAction()

	shouldWhitelist := ctrl.authWhitelist[action]

	if !shouldWhitelist {
		if err != nil {
			ctrl.Data["json"] = map[string]interface{}{"message": "You are unauthorised"}
			ctrl.ServeJSON()
			ctrl.StopRun()
		} else {
			user := models.User{Id: tokenInfo.GetUserID()}

			if err = models.Orm.Read(&user); err != nil {
				ctrl.Data["json"] = map[string]interface{}{"message": "You are unauthorised"}
				ctrl.ServeJSON()
				ctrl.StopRun()
			} else {
				ctrl.user = &user
			}
		}
	}
}
