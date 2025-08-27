package gumble

import (
	"crypto/tls"

	"github.com/bytecntrl/gumble/internal/connection"
	"github.com/bytecntrl/gumble/internal/model"
)

// NewMumble creates a new Mumble client instance using the provided configuration and TLS settings.
//
// It establishes a TCP connection to the Mumble server using the specified network parameters
// and wraps it in a Mumble client object.
//
// Parameters:
//   - config: Pointer to a Config struct containing the server address and authentication credentials.
//   - tlsConfig: Pointer to a tls.Config used for the TLS handshake. This can include settings like
//     certificates, root CAs, or InsecureSkipVerify for development purposes.
//
// Returns:
//   - A pointer to a new Mumble instance on success.
//   - An error if the TCP connection fails or TLS negotiation encounters a problem.
func NewMumble(config *Config, tlsConfig *tls.Config) (*Mumble, error) {
	conn, err := connection.NewTCPConn(
		&model.NetConfig{
			Hostname: config.Hostname,
			Port:     config.Port,
		},
		tlsConfig,
	)

	if err != nil {
		return nil, err
	}

	return &Mumble{
		Config:  config,
		TCPConn: conn,
	}, nil
}
