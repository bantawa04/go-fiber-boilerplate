package validator

import (
	"fmt"

	"github.com/bantawao4/gofiber-boilerplate/app/response"
	"github.com/go-playground/validator/v10"
)

type TodoValidator struct {
	*validator.Validate
}

func NewTodoValidator() TodoValidator {
	v := validator.New()
	_ = v.RegisterValidation("status", func(fl validator.FieldLevel) bool {
		if fl.Field().String() != "" {
			validStatuses := map[string]bool{
				"pending":     true,
				"in_progress": true,
				"completed":   true,
			}
			return validStatuses[fl.Field().String()]
		}
		return true
	})
	return TodoValidator{
		Validate: v,
	}
}

func (cv TodoValidator) generateValidationMessage(field string, rule string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	case "status":
		return fmt.Sprintf("Field '%s' must be one of: pending, in_progress, completed", field)
	case "min":
		return fmt.Sprintf("Field '%s' is too short", field)
	case "max":
		return fmt.Sprintf("Field '%s' is too long", field)
	default:
		return fmt.Sprintf("Field '%s' is not valid.", field)
	}
}

func (cv TodoValidator) GenerateValidationResponse(err error) []response.ValidationError {
	var validations []response.ValidationError
	for _, value := range err.(validator.ValidationErrors) {
		field, rule := value.Field(), value.Tag()
		validation := response.ValidationError{Field: field, Message: cv.generateValidationMessage(field, rule)}
		validations = append(validations, validation)
	}
	return validations
}
