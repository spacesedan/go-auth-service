package apperrors

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func AlreadyExists() error {
	return errors.New("user already exists")
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required" + fe.Param()
	case "min":
		return "This field should be greater than " + fe.Param()
	default:
		return "Unknown error"
	}
}
