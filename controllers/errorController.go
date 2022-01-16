package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}

func (ctrl *ErrorController) ErrorDb() {
	fmt.Println("Error Db")
	fmt.Println(ctrl.Data)
	ctrl.Data["json"] = "database is now down"
}

func (ctrl *ErrorController) Render() error {
	ctrl.ServeJSON()
	return nil
}
