package controllers

import (
	beego "github.com/beego/beego/v2/server/web"

	"docket-beego/auth"
)

type OauthController struct {
	beego.Controller
}

// func (ctrl OauthController) TokenHandler(c *context.Context) {
// 	err := utils.GetSrv().HandleTokenRequest(c.Writer, c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Wrong password",
// 		})
// 	}
// }

// @router /token [post]
func (u *OauthController) CreateToken() (map[string]interface{}, error) {
	err := auth.Srv.HandleTokenRequest(u.Ctx.ResponseWriter, u.Ctx.Request)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
