package models

type ErrorQuery struct {
	message string
	code    string
	err     error
}

type errorCode = string

type errorEnum struct {
	Internal errorCode
	Invalid  errorCode
}

var ErrorEnum = &errorEnum{
	Internal: "internal",
	Invalid:  "invalid",
}

func NewErrorQuery(message string, code errorCode, err error) *ErrorQuery {
	return &ErrorQuery{message, code, err}
}

func (eq *ErrorQuery) Error() string {
	return eq.message
}

func (eq *ErrorQuery) Unwrap() error {
	return eq.err
}

// Code return error code based on ErrorEnum
func (eq *ErrorQuery) Code() errorCode {
	return eq.code
}
