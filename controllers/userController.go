package controllers

import (
	"docket-beego/models"
	"docket-beego/utils"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web/pagination"
	"github.com/google/uuid"
)

type RegisterUserBody struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
	Password  *string `json:"password"`
}

func (b *RegisterUserBody) validate() ([]*validation.Error, error) {
	valid := validation.Validation{}
	valid.Required(b.FirstName, "first_name")
	valid.Required(b.LastName, "last_name")
	valid.Required(b.Username, "username")
	valid.Required(b.Password, "password")

	ok, err := valid.Valid(b)

	if err != nil {
		return nil, err
	} else {
		if !ok {
			return valid.Errors, nil
		}

		return nil, nil
	}
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
	errs, _ := newUserBody.validate()

	if len(errs) > 0 {
		// responseErr := utils.GetErrorResponse(*ctrl.Ctx, utils.GetValidationError(), nil)
		// return &responseErr
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
	// ctrl.Ctx.Re

	pgF := utils.NewPaginationQuery()

	if err := ctrl.ParseForm(&pgF); err != nil {
		logs.Error(err)
	} else {
		logs.Info("No err")
	}

	paginatedQs, _, err := utils.GetPagninatedQs(pgF, ctrl.Ctx.Request, qs)

	if err != nil {
		logs.Error(err)
	}

	result := []models.User{}

	_, err = paginatedQs.All(&result)

	if err != nil {
		logs.Error(err)
	}

	pagination.SetPaginator(ctrl.Ctx, 5, 10)

	response := utils.GetSuccessResponse(*ctrl.Ctx, result, 201, nil)

	return &response
}
