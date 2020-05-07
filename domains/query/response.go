package query

import "github.com/tamaApotek/tama-go-server/domains/apperror"

type Response struct {
	Message string             `json:"message"`
	Code    apperror.ErrorCode `json:"code,omitempty"`
	Data    interface{}        `json:"data,omitempty"`
}
