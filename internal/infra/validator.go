package infra

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// InitValidator inisialisasi global validator
func InitValidator() {
	validate = validator.New()
}

// ValidateStruct memvalidasi struct dan mengembalikan map field -> error message
func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		tag := e.Tag()
		param := e.Param()

		errors[field] = getErrorMessage(field, tag, param)
	}
	return errors
}

// getErrorMessage menghasilkan pesan error yang ramah
func getErrorMessage(field, tag, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, param)
	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", field, param)
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", field, param)
	default:
		return fmt.Sprintf("%s is invalid (%s)", field, tag)
	}
}
