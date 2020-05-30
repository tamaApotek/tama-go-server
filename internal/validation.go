package internal

import (
	"github.com/go-playground/validator/v10"
	"github.com/tamaApotek/tama-go-server/internal/validation"
)

func InitValidation(v *validator.Validate) {
	validation.ValidateDateStrTag(v)
}
