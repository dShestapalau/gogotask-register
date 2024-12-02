package utils

import (
	"errors"
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

func HandleError(err error) ErrorResponse {
	response := ErrorResponse{}

	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		out := make([]ApiError, len(validationErrors))

		for i, fe := range validationErrors {
			out[i] = ApiError{fe.Field(), msgToError(fe)}
		}

		response.Errors = out
	}

	return response
}

func msgToError(tag validator.FieldError) string {
	switch tag.Tag() {
	case "required":
		return "field " + tag.Field() + " is required"
	case "notEmpty":
		return "field " + tag.Field() + " is empty"
	case "min":
		return "field " + tag.Field() + " should be greater than " + tag.Param() + " symbols"
	case "max":
		return "field " + tag.Field() + " should be less than " + tag.Param() + " symbols"
	}

	return ""
}
