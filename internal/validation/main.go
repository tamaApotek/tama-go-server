package validation

import "github.com/go-playground/validator/v10"

var valmap map[string]validator.Func

func Init(v *validator.Validate) {
	validateRoleTag(v)
	validateDateStrTag(v)
}
