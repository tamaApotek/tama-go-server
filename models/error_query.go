package models

type errorQuery struct {
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

func NewErrorQuery(message string, code errorCode, err error) *errorQuery {
	return &errorQuery{message, code, err}
}

func (eq *errorQuery) Error() string {
	return eq.message
}

func (eq *errorQuery) Unwrap() error {
	return eq.err
}
