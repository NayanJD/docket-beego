package controllers

import (
	"fmt"
)

// var TaskControllerInstance = TaskController{au}

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
// @router /task [get]
func (ctrl *TaskController) GetTask(apiVersion *string) (*Dummy, error) {
	fmt.Println("From GetTask")
	ctrl.Ctx.Output.SetStatus(201)
	return &Dummy{Id: "dfadf"}, nil
}

// func (ctrl *TaskController) Finish() {
// 	fmt.Println(ctrl.)
// }
