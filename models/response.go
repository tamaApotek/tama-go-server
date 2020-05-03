package models

import "github.com/tamaApotek/tama-go-server/query"

type Response struct {
	Message string          `json:"message"`
	Code    query.ErrorCode `json:"code,omitempty"`
	Data    interface{}     `json:"data,omitempty"`
}
