package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/tamaApotek/tama-go-server/domain/role"
)

// const role = "role"

var roleValidation validator.Func = func(fl validator.FieldLevel) bool {
	r, ok := fl.Field().Interface().(role.Role)
	if !ok {
		return false
	}

	return r.IsValid()
}

func validateRoleTag(v *validator.Validate) {
	v.RegisterValidation("role", roleValidation)
}
