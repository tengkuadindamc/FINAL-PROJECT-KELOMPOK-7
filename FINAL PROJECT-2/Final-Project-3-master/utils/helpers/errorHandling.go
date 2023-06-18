package helpers

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func msgForTag(tag string, field string) string {
	if tag == "required" {
		return  field + " is required"
	} else if tag == "email" {
		return "The field must be a email"
	} else if tag == "min" {
		return field + " must be more than 6 character"
	} else {
		return "Invalid " + field
	}
}

func FormatError(err error) []string {
	var errorr []string
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, e := range err.(validator.ValidationErrors) {
			errorr = append(errorr, msgForTag(e.Tag(), e.Field()))
		}
	}
	return errorr
}
