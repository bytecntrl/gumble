package gumble

import "github.com/bytecntrl/gumble/internal/model"

// NewConfig creates and returns a new Config instance for connecting to a Mumble server.
//
// It initializes the embedded NetConfig and AuthConfig structs with the provided
// hostname, port, username, and password. This function is typically used to simplify
// configuration setup before creating a Mumble client.
//
// Parameters:
//   - Hostname: The IP address or domain name of the Mumble server.
//   - Port: The port number to connect to (default Mumble port is 64738).
//   - Username: The username to authenticate with.
//   - Password: The password to authenticate with (can be empty if server allows anonymous login).
//
// Returns:
//   - A pointer to a fully initialized Config struct.
func NewConfig(Hostname string, Port int, Username string, Password string) *Config {
	return &Config{
		ConnectionConfig: model.ConnectionConfig{
			NetConfig: model.NetConfig{
				Hostname: Hostname,
				Port:     Port,
			},
			AuthConfig: model.AuthConfig{
				Username: Username,
				Password: Password,
			},
		},
	}
}
