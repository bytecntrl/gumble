package errs

// NewTCPConnectionError creates a new BaseError indicating a failure to establish or maintain a TCP connection.
// It wraps the underlying error and uses the CodeConnection error code.
func NewTCPConnectionError(message string, err error) *BaseError {
	return NewBaseError(CodeConnection, message, err)
}

// NewTCPWriteError creates a new BaseError for a failure during a write operation over a TCP connection.
// It wraps the underlying error and uses the CodeConnection error code.
func NewTCPWriteError(message string, err error) *BaseError {
	return NewBaseError(CodeConnection, message, err)
}

// NewTCPIncompleteWrite creates a new BaseError indicating that a TCP write operation completed only partially.
// It uses the CodeConnection error code and does not wrap any underlying error.
func NewTCPIncompleteWrite(message string) *BaseError {
	return NewBaseError(CodeConnection, message, nil)
}
