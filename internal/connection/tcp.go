package connection

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"

	"github.com/bytecntrl/gumble/errs"
	"github.com/bytecntrl/gumble/internal/model"
)

// TCPConn is a wrapper around a TLS connection.
type TCPConn struct {
	sync.Mutex

	conn *tls.Conn // Secure TCP connection
}

// NewTCPConn creates a new TCPConn using the connection settings.
func NewTCPConn(cfg *model.NetConfig, tlsConfig *tls.Config) (*TCPConn, error) {
	dialer := newDialer()
	address := fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port)

	// Connect to server with TLS
	conn, err := tls.DialWithDialer(dialer, "tcp", address, tlsConfig)
	if err != nil {
		return nil, errs.NewTCPConnectionError("failed to connect to server", err)
	}

	return &TCPConn{
		conn: conn,
	}, nil
}

// Close the connection with the server
func (conn *TCPConn) Close() error {
	if err := conn.conn.Close(); err != nil {
		return errs.NewTCPConnectionError("failed to close TCP connection", err)
	}
	return nil
}

// CreateHeader builds a 6-byte header for a Mumble protocol message.
//
// The header format is:
// - First 2 bytes: the message type (uint16, big-endian)
// - Next 4 bytes: the payload length (uint32, big-endian)
//
// This header must be sent before the protobuf-encoded message.
func (conn *TCPConn) CreateHeader(messageTypeVersion uint16, payloadLength uint32) []byte {
	header := make([]byte, 6)

	binary.BigEndian.PutUint16(header[0:2], messageTypeVersion)
	binary.BigEndian.PutUint32(header[2:6], payloadLength)

	return header
}

// WritePacket serializes and writes a proto.Message to the TCP connection.
// It checks if the message is allowed, builds the header, and sends the packet.
func (conn *TCPConn) WritePacket(msg proto.Message) error {
	conn.Lock()
	defer conn.Unlock()

	messageType, ok := GetTCPMessageType(msg)
	if !ok {
		return errs.NewInvalidMessageType("unsupported or invalid TCP message type")
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		return errs.NewSerializationError("failed to marshal proto", err)
	}

	header := conn.CreateHeader(messageType, uint32(len(data)))
	packet := append(header, data...)

	n, err := conn.conn.Write(packet)
	if err != nil {
		return errs.NewTCPWriteError("failed to write to TCP connection", err)
	}
	if n != len(packet) {
		return errs.NewTCPIncompleteWrite("incomplete TCP packet write")
	}

	return nil
}
