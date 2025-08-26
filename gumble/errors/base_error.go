package errors

import "fmt"

type BaseError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *BaseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *BaseError) Unwrap() error {
	return e.Err
}

func NewBaseError(code ErrorCode, message string, err error) *BaseError {
	return &BaseError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
