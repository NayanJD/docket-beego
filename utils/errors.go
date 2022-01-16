package utils

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/core/validation"
)

func GetValidationError(errs []*validation.Error) ApiError {
	errorMessages := []string{}
	for _, err := range errs {
		fmt.Println(err.Key, err.Message)
		errorMessages = append(errorMessages, err.Message+": "+err.Key)
	}

	return ApiError{
		statusCode: http.StatusUnprocessableEntity,
		ResponseError: ResponseError{
			ErrorCode:     "VALIDATION_ERROR",
			ErrorMessages: errorMessages,
		},
	}
}

func GetConflictError(key string, value string) ApiError {
	return ApiError{
		statusCode: http.StatusConflict,
		ResponseError: ResponseError{
			ErrorCode:     "CONFLICT",
			ErrorMessages: []string{key + " with " + value + " already exists."},
		},
	}
}
