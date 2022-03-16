package apperrors

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var (
	AlreadyExistsErr   = errors.New("user already exists")
	NoUserErr          = errors.New("no user found")
	WrongPasswordErr   = errors.New("wrong password")
	GeneratingTokenErr = errors.New("token not generated")
	NotValidTokenErr   = errors.New("token not valid")
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
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
