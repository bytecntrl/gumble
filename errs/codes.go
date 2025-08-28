package errs

// ErrorCode represents a machine-readable code that categorizes the type of error.
// It allows errors to be programmatically identified and handled.
type ErrorCode string

const (
	// CodeSerialization indicates an error occurred during serialization or deserialization
	// of a message, such as a failure in encoding or decoding data formats.
	CodeSerialization ErrorCode = "SERIALIZATION_ERROR"

	// CodeConnection indicates a failure related to network connections,
	// including establishing, maintaining, or writing to a connection.
	CodeConnection ErrorCode = "CONNECTION_ERROR"

	// CodeInvalidMessage indicates that an invalid, unexpected, or unsupported
	// message type or format was encountered during processing.
	CodeInvalidMessage ErrorCode = "INVALID_MESSAGE"
)
