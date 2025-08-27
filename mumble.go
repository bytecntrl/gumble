package gumble

import (
	"github.com/bytecntrl/gumble/internal/connection"
)

// Mumble is the struct used for the connection this the server
type Mumble struct {
	Config  *Config             // Config contains the client configurations
	TCPConn *connection.TCPConn // TCPConn contains the methods for the connection
}
