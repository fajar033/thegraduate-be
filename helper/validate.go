package helper

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

var validate *validator.Validate

func Validate[t any](data t) []map[string]any {

	validate = validator.New()

	err := validate.Struct(&data)

	if err != nil {

		var message []map[string]interface{}
		for _, e := range err.(validator.ValidationErrors) {

			message = append(message, map[string]interface{}{
				"field":   e.Field(),
				"message": customMessage(e),
			})

		}

		return message
	}
	return nil
}

func customMessage(e validator.FieldError) string {

	switch e.Tag() {
	case "required":
		return "This field is required" + e.Param()
	case "min":
		return "length of this field should be more than " + e.Param() + " characters"
	case "max":
		return "length of this field should be less than " + e.Param() + " characters"
	case "alpha":
		return "Must be alphabet only"
	case "email":
		return "Email is not valid, please check again"
	case "oneof":
		allowedOptions := strings.Replace(e.Param(), " ", " or ", -1)
		return "Value must be one of the allowed options: " + allowedOptions
	}

	return ""
}
