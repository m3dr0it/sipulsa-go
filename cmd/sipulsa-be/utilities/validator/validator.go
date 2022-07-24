package validator

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type PayloadValidatorI interface {
	getValidator() validator.Validate
}

type payloadValidator struct {
	Validator *validator.Validate
}

func (cv *payloadValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func (v *payloadValidator) getValidator() *validator.Validate {
	return v.Validator
}

func SetValidator(e *echo.Echo) {
	e.Validator = &payloadValidator{
		Validator: validator.New(),
	}
}
