package query

type Response struct {
	Message string      `json:"message"`
	Code    ErrorCode   `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
