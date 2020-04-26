package models

type Response struct {
	Message string      `json:"message"`
	Code    errorCode   `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
