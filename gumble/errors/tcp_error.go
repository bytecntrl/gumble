package errors

func NewTCPConnectionError(message string, err error) *BaseError {
	return NewBaseError(CodeConnection, message, err)
}

func NewTCPWriteError(message string, err error) *BaseError {
	return NewBaseError(CodeConnection, message, err)
}

func NewTCPIncompleteWrite(message string) *BaseError {
	return NewBaseError(CodeConnection, message, nil)
}
