package validation

import (
	"errors"
	"fmt"
	errorUtils "github.com/Yutsss/FP-PBKK-GOLANG/BE/utility/error"
	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return CustomErrorMessage(validationErrors[0])
		} else {
			return errorUtils.ErrInternalServer
		}
	}

	return nil
}

func CustomErrorMessage(err validator.FieldError) error {
	switch err.Tag() {
	case "required":
		return fmt.Errorf("%s is required", err.Field())
	case "min":
		return fmt.Errorf("%s cannot be under %s characters", err.Field(), err.Param())
	case "email":
		return fmt.Errorf("Invalid email format")
	default:
		return fmt.Errorf("Invalid value for %s", err.Field())
	}
}
