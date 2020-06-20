package bootstrap

import "github.com/go-playground/validator/v10"

type App struct {
	Validator *validator.Validate
	Handler   *Handler
}
