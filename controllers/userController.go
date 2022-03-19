package controllers

import (
	"docket-beego/models"
	"docket-beego/utils"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/google/uuid"
)

type RegisterUserBody struct {
	FirstName *string `json:"first_name" validate:"required"`
	LastName  *string `json:"last_name" validate:"required"`
	Username  *string `json:"username" validate:"required"`
	Password  *string `json:"password" validate:"required"`
}

type UserController struct {
	AuthController
}

func NewUserController() *UserController {
	return &UserController{AuthController: AuthController{authWhitelist: map[string]bool{"Register": true}}}
}

// @Param   apiVersion     query   string false "v1"       ""
// @Param	newUserBody	body	{RegisterUserBody}	true	"New user object"
// @router /register [post]
func (ctrl *UserController) Register(apiVersion *string, newUserBody *RegisterUserBody) *utils.GenericResponse {
	newId := uuid.New().String()
	isSuperuser := false
	isStaff := false
	verr, _ := utils.ValidatePayload(newUserBody)

	if verr != nil {
		responseErr := utils.GetErrorResponse(*ctrl.Ctx, utils.GetValidationError(verr), nil)
		return &responseErr
	}

	doesUserExist := models.Orm.QueryTable("users").Filter("username", *newUserBody.Username).Exist()

	if doesUserExist {
		responseErr := utils.GetErrorResponse(*ctrl.Ctx, utils.GetConflictError("username", *newUserBody.Username), nil)
		return &responseErr
	}

	newUser := models.User{
		Id:          newId,
		FirstName:   newUserBody.FirstName,
		LastName:    newUserBody.LastName,
		Username:    newUserBody.Username,
		Password:    newUserBody.Password,
		IsSuperuser: &isSuperuser,
		IsStaff:     &isStaff,
	}

	_, err := models.Orm.Insert(&newUser)

	if err != nil {
		fmt.Println(err)
	}

	newUser.SetPassword(newUser.Password)

	_, err = models.Orm.Update(&newUser)

	response := utils.GetSuccessResponse(*ctrl.Ctx, newUser, 201, nil)

	return &response
}

// @Param   apiVersion     query   string false "v1"
// @router	/	[get]
func (ctrl *UserController) GetUsers(apiVersion *string) *utils.GenericResponse {

	qs := models.Orm.QueryTable(models.User{})

	paginatedQs, err := utils.GetPagninatedQs(ctrl.Controller, qs)

	if err != nil {
		logs.Error(err)
		responseErr := utils.GetErrorResponse(*ctrl.Ctx, utils.InternalError, nil)
		return &responseErr
	}

	result := []models.User{}

	_, err = paginatedQs.All(&result)

	if err != nil {
		logs.Error(err)
		responseErr := utils.GetErrorResponse(*ctrl.Ctx, utils.InternalError, nil)
		return &responseErr
	}

	response := utils.GetSuccessResponse(*ctrl.Ctx, result, 201, nil)

	return &response
}
