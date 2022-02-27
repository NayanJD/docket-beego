package utils

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorToText(e validator.FieldError) string {
	lowerCaseField := strings.ToLower(e.Field())
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", lowerCaseField)
	case "max":
		return fmt.Sprintf(
			"%s cannot be longer than %s",
			lowerCaseField,
			e.Param(),
		)
	case "min":
		return fmt.Sprintf(
			"%s must be longer than %s",
			lowerCaseField,
			e.Param(),
		)
	case "email":
		return "Invalid email format"
	case "len":
		return fmt.Sprintf(
			"%s must be %s characters long",
			lowerCaseField,
			e.Param(),
		)
	}
	return fmt.Sprintf("%s is not valid", lowerCaseField)
}

func GetValidationError(errs *validator.ValidationErrors) ApiError {
	errorMessages := []string{}
	for _, err := range *errs {
		errorMessages = append(errorMessages, ValidationErrorToText(err))
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
