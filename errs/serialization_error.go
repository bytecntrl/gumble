package errs

// NewSerializationError creates a new BaseError indicating a serialization or encoding/decoding failure.
// It wraps the underlying error and uses the CodeSerialization error code.
func NewSerializationError(message string, err error) *BaseError {
	return NewBaseError(CodeSerialization, message, err)
}

// NewInvalidMessageType creates a new BaseError for situations where an unexpected or unsupported message type is encountered.
// It uses the CodeInvalidMessage error code and does not wrap any underlying error.
func NewInvalidMessageType(message string) *BaseError {
	return NewBaseError(CodeInvalidMessage, message, nil)
}
