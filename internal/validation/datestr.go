package validation

import (
	"time"

	"github.com/go-playground/validator/v10"
)

const datestrLayout = "2006-01-02"

var datestrValidation validator.Func = func(fl validator.FieldLevel) bool {
	datestr, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	_, err := time.Parse(datestrLayout, datestr)
	if err != nil {
		return false
	}

	return true

}

// ValidateDateStrTag validate date with YYYY-MM-DD format
func ValidateDateStrTag(v *validator.Validate) {
	v.RegisterValidation("datestr", datestrValidation)
}
