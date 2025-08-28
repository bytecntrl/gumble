package errs

import "fmt"

// BaseError represents a structured error with a code, message, and optional wrapped error.
// It implements the error interface and supports error unwrapping for use with errors.Is and errors.As.
type BaseError struct {
	Code    ErrorCode
	Message string
	Err     error
}

// Error returns the string representation of the BaseError.
// It includes the error code, message, and any wrapped error.
func (e *BaseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error, if any.
// This allows BaseError to be compatible with errors.Is and errors.As.
func (e *BaseError) Unwrap() error {
	return e.Err
}

// NewBaseError creates a new BaseError with the given code, message, and optional underlying error.
func NewBaseError(code ErrorCode, message string, err error) *BaseError {
	return &BaseError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
