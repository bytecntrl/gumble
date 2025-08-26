package model

// NetConfig holds the network address (host and port) of the server.
type NetConfig struct {
	Hostname string // Server address (IP or domain)
	Port     int    // Port number
}

// AuthConfig holds the authentication credentials.
type AuthConfig struct {
	Username string // Username for login
	Password string // Password for login
}

// ConnectionConfig contains both network and authentication settings.
type ConnectionConfig struct {
	NetConfig  // Embedded network configuration
	AuthConfig // Embedded authentication configuration
}
