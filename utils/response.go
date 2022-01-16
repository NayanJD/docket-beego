package utils

import (
	"net/http"

	"github.com/beego/beego/v2/server/web/context"
)

type ResponseError struct {
	ErrorCode     string   `json:"error_code"`
	ErrorMessages []string `json:"error_messages"`
}

type GenericResponse struct {
	Data      *interface{}   `json:"data"`
	Errors    *ResponseError `json:"errors"`
	IsSuccess *bool          `json:"is_success"`
	Meta      *interface{}   `json:"meta"`
}
type ApiError struct {
	statusCode int
	ResponseError
}

func IsStatusSuccess(code int) bool {
	return code >= http.StatusOK && code < http.StatusMultipleChoices
}

func GetSuccessResponse(context context.Context, data interface{}, status int, meta *interface{}) GenericResponse {
	// context.Output.SetStatus(status)

	isSuccess := IsStatusSuccess(status)

	return GenericResponse{
		Data:      &data,
		Errors:    nil,
		IsSuccess: &isSuccess,
		Meta:      meta,
	}
}

func GetErrorResponse(context context.Context, err ApiError, meta *interface{}) GenericResponse {
	context.Output.SetStatus(err.statusCode)

	isSuccess := IsStatusSuccess(err.statusCode)

	return GenericResponse{
		Data:      nil,
		Errors:    &err.ResponseError,
		IsSuccess: &isSuccess,
		Meta:      meta,
	}
}
