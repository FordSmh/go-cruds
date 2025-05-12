package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var Validate = validator.New()

func PrepareErrorMessage(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			logrus.Info(e.Field())
			field := e.Field()
			switch e.Tag() {
			case "required":
				errors[field] = field + " is required"
			case "email":
				errors[field] = "Invalid email format"
			case "min":
				errors[field] = field + " must be at least " + e.Param() + " characters long"
			case "max":
				errors[field] = field + " must be at most " + e.Param() + " characters long"
			default:
				errors[field] = "Invalid " + field
			}
		}
	}
	return errors
}
