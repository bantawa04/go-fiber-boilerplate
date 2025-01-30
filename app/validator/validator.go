package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Validator struct {
	*validator.Validate
}

func NewValidator() Validator {
	v := validator.New()

	// Register custom validations
	_ = v.RegisterValidation("phone", validatePhone)
	_ = v.RegisterValidation("gender", validateGender)
	_ = v.RegisterValidation("email", validateEmail)

	return Validator{
		Validate: v,
	}
}

func validatePhone(fl validator.FieldLevel) bool {
	if fl.Field().String() != "" {
		match, _ := regexp.MatchString("^[- +()]*[0-9][- +()0-9]*$", fl.Field().String())
		return match
	}
	return true
}

func validateEmail(fl validator.FieldLevel) bool {
	if fl.Field().String() != "" {
		match, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, fl.Field().String())
		return match
	}
	return true
}

func validateGender(fl validator.FieldLevel) bool {
	if fl.Field().String() != "" {
		gender := fl.Field().String()
		return gender == "male" || gender == "female" || gender == "other"
	}
	return true
}

func (v Validator) GenerateValidationMessage(field string, rule string) string {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is required", field)
	case "phone":
		return fmt.Sprintf("Field '%s' is not a valid phone number", field)
	case "gender":
		return fmt.Sprintf("Field '%s' must be 'male', 'female', or 'other'", field)
	case "email":
		return fmt.Sprintf("Field '%s' is not a valid email address", field)
	default:
		return fmt.Sprintf("Field '%s' is not valid", field)
	}
}

func (v Validator) ValidateStruct(data interface{}) []ValidationError {
	var validationErrors []ValidationError

	err := v.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Field:   err.Field(),
				Message: v.GenerateValidationMessage(err.Field(), err.Tag()),
			})
		}
	}
	return validationErrors
}
