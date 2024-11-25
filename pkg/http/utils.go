package utils

import (
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidationRules() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notEmpty", notEmpty)
	}

}

var notEmpty validator.Func = func(fl validator.FieldLevel) bool {
	field, ok := fl.Field().Interface().(string)

	if ok {
		return strings.TrimSpace(field) != ""
	}
	return true
}
