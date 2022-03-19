package controllers

import (
	"docket-beego/models"
	"docket-beego/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/pagination"
	"github.com/google/uuid"
)

// var TaskControllerInstance = TaskController{au}

type CreateTaskBody struct {
	Description string `json:"description" validate:"required"`
	ScheduledAt string `json:"scheduled_at" validate:"required,datetime=2006-01-02T15:04:05Z|datetime=2006-01-02T15:04:05+0000"`
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

	verr, _ := utils.ValidatePayload(newTaskBody)

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

	qs := models.Orm.QueryTable(models.Task{})

	paginatedQs, err := utils.GetPagninatedQs(ctrl.Controller, qs)

	if err != nil {
		logs.Error(err)
	}

	result := []models.Task{}

	_, err = paginatedQs.Filter("user_id", ctrl.user.Id).All(&result)

	if err != nil {
		logs.Error(err)
	}

	pagination.SetPaginator(ctrl.Ctx, 5, 10)

	response := utils.GetSuccessResponse(*ctrl.Ctx, result, 200, nil)

	return &response
}
