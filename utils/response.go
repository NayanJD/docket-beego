package utils

import (
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

type ResponseError struct {
	ErrorCode     string   `json:"error_code"`
	ErrorMessages []string `json:"error_messages"`
}

type GenericResponse struct {
	Data      *interface{}            `json:"data"`
	Errors    *ResponseError          `json:"errors"`
	IsSuccess *bool                   `json:"is_success"`
	Meta      *map[string]interface{} `json:"meta"`
}

type ApiError struct {
	statusCode int
	ResponseError
}

var InternalError = ApiError{
	statusCode: http.StatusInternalServerError,
	ResponseError: ResponseError{
		ErrorCode:     "INTERNAL_SERVER_ERROR",
		ErrorMessages: []string{"Something went wrong"},
	},
}

func IsStatusSuccess(code int) bool {
	return code >= http.StatusOK && code < http.StatusMultipleChoices
}

func GetSuccessResponse(context context.Context, data interface{}, status int, meta interface{}) GenericResponse {
	// context.Output.SetStatus(status)

	isSuccess := IsStatusSuccess(status)

	paginationData := context.Input.GetData(PAGINATION_DATA_KEY)

	metaMap, ok := meta.(map[string]interface{})

	if !ok {
		metaMap = map[string]interface{}{"pagination": paginationData}
	} else {
		metaMap = map[string]interface{}{}
	}

	return GenericResponse{
		Data:      &data,
		Errors:    nil,
		IsSuccess: &isSuccess,
		Meta:      &metaMap,
	}
}

func GetErrorResponse(context context.Context, err ApiError, meta interface{}) GenericResponse {
	context.Output.SetStatus(err.statusCode)

	isSuccess := IsStatusSuccess(err.statusCode)

	paginationData := context.Input.GetData(PAGINATION_DATA_KEY)

	metaMap, ok := meta.(map[string]interface{})

	if !ok {
		metaMap = map[string]interface{}{"pagination": paginationData}
	}

	return GenericResponse{
		Data:      nil,
		Errors:    &err.ResponseError,
		IsSuccess: &isSuccess,
		Meta:      &metaMap,
	}
}

func RecoverPanicFunc(ctx *context.Context, cfg *beego.Config) {
	if err := recover(); err != nil {
		if err == beego.ErrAbort {
			return
		}
		if !cfg.RecoverPanic {
			panic(err)
		}

		isSuccess := false

		ctx.Output.SetStatus(http.StatusInternalServerError)
		ctx.Output.JSON(GenericResponse{
			Data:      nil,
			Errors:    &InternalError.ResponseError,
			IsSuccess: &isSuccess,
			Meta:      nil,
		}, true, false)

	}
}
