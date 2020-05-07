package apperror

type ErrorCode string

type AppError struct {
	message string
	code    ErrorCode
	err     error
}

const (
	// ErrInternal represent internal server error
	ErrInternal ErrorCode = "internal"
	// ErrInvalid represent client data error
	ErrInvalid ErrorCode = "invalid"
)

func (ec *ErrorCode) String() string {
	return string(*ec)
}

// New construct new apperror
func New(message string, code ErrorCode, err error) error {
	return &AppError{message, code, err}
}

func (ae *AppError) Error() string {
	return ae.message
}

func (ae *AppError) Unwrap() error {
	return ae.err
}

func (ae *AppError) Code() ErrorCode {
	return ae.code
}
