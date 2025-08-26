package errors

type ErrorCode string

const (
	CodeSerialization  ErrorCode = "SERIALIZATION_ERROR"
	CodeConnection     ErrorCode = "CONNECTION_ERROR"
	CodeInvalidMessage ErrorCode = "INVALID_MESSAGE"
)
