package query

type ErrorCode string

type ErrorQuery struct {
	message string
	code    ErrorCode
	err     error
}

type errorEnum struct {
	Internal ErrorCode
	Invalid  ErrorCode
}

var ErrorEnum = &errorEnum{
	Internal: "internal",
	Invalid:  "invalid",
}

func NewErrorQuery(message string, code ErrorCode, err error) *ErrorQuery {
	return &ErrorQuery{message, code, err}
}

func (eq *ErrorQuery) Error() string {
	return eq.message
}

func (eq *ErrorQuery) Unwrap() error {
	return eq.err
}

// Code return query code based on ErrorEnum
func (eq *ErrorQuery) Code() ErrorCode {
	return eq.code
}
