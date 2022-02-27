package controllers

import (
	"fmt"

	"docket-beego/models"
	"docket-beego/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// var TaskControllerInstance = TaskController{au}

type CreateTaskBody struct {
	Description string `json:"description" validate:"required"`
	ScheduledAt string `validate:"required,datetime=2006-01-02T15:04:05Z|datetime=2006-01-02T15:04:05+0000" json:"scheduled_at"`
}

func (b *CreateTaskBody) validate() (*validator.ValidationErrors, error) {
	err := utils.Validator.Struct(b)

	switch v := err.(type) {
	case *validator.InvalidValidationError:
		panic("InvalidValidationError thrown")
	case validator.ValidationErrors:
		return &v, nil
	default:
		return nil, nil
	}
}

type TaskController struct {
	AuthController
}

type Dummy struct {
	Id string `json:"id"`
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (ctrl *TaskController) New() TaskController {
	return TaskController{AuthController: ctrl.AuthController}
}

// @Param   apiVersion     query   string false "v1"       ""
// @Param	newTaskBody	body	{CreateTaskBody}	true	"New task object"
// @router	/	[post]
func (ctrl *TaskController) CreateTask(apiVersion *string, newTaskBody *CreateTaskBody) *utils.GenericResponse {

	verr, _ := newTaskBody.validate()

	logs.Error(verr)

	if verr != nil {
		responseErr := utils.GetErrorResponse(*ctrl.Ctx, utils.GetValidationError(verr), nil)
		return &responseErr
	}

	newId := uuid.New().String()
	parsedScheduledAt, err := utils.ParseTime(newTaskBody.ScheduledAt)

	if err != nil {
		panic("ScheduledAt was passed by validators but not parsed by utils.ParseTime")
	}

	newTask := models.Task{
		Id:          newId,
		Description: &newTaskBody.Description,
		ScheduledAt: parsedScheduledAt,
		User: &models.User{
			Id: ctrl.user.Id,
		},
	}

	_, err = models.Orm.Insert(&newTask)

	if err != nil {
		logs.Error(err)
	}

	response := utils.GetSuccessResponse(*ctrl.Ctx, newTask, 200, nil)

	return &response
}

// @Param   apiVersion     query   string false "v1"       ""
// @router	/	[get]
func (ctrl *TaskController) GetTask(apiVersion *string) *utils.GenericResponse {

	result := []models.Task{}

	_, err := models.Orm.QueryTable("tasks").Filter("user_id", ctrl.user.Id).All(&result)

	// for _, task := range result {
	// 	fmt.Println(task.User)
	// }

	if err != nil {
		fmt.Println(err)
	}

	response := utils.GetSuccessResponse(*ctrl.Ctx, result, 200, nil)

	return &response
}
