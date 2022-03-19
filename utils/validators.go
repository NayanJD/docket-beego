package utils

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() {
	Validator = validator.New()

	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

// 2006-01-02T15:04:05Z
// 2006-01-02T15:04:05+07:00
func ParseTime(timeString string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, timeString)

	if err != nil {
		t, err := time.Parse("2006-01-02T15:04:05+0000", timeString)

		if err != nil {
			return nil, err
		}

		return &t, err
	} else {
		return &t, nil
	}
}

func ValidatePayload(b interface{}) (*validator.ValidationErrors, error) {
	err := Validator.Struct(b)

	switch v := err.(type) {
	case *validator.InvalidValidationError:
		panic("InvalidValidationError thrown")
	case validator.ValidationErrors:
		return &v, nil
	default:
		return nil, nil
	}
}
