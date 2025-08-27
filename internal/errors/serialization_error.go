package errors

func NewSerializationError(message string, err error) *BaseError {
	return NewBaseError(CodeSerialization, message, err)
}

func NewInvalidMessageType(message string) *BaseError {
	return NewBaseError(CodeInvalidMessage, message, nil)
}
