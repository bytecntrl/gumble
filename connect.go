package gumble

import (
	"runtime"

	"github.com/bytecntrl/gumble/internal"
)

// Connect establishes a connection to the Mumble server.
//
// It sends the Version packet with client version and platform information,
// followed by the Authenticate packet using the configured credentials.
// If either step fails, it returns the corresponding error.
func (m *Mumble) Connect() error {
	// Generate the VersionV2 number from major, minor, and patch values.
	version := internal.GetVersionV2(MumbleVersionMajor, MumbleVersionMinor, MumbleVersionPatch)

	// Send the Version packet to the server with relevant client details.
	if err := m.Version(&VersionOptions{
		VersionV2: version,
		Release:   MumbleRelease,
		Os:        runtime.GOOS,
		OsVersion: runtime.GOARCH,
	}); err != nil {
		return err
	}

	// Send the Authenticate packet using configuration values.
	if err := m.Authenticate(&AuthenticateOptions{
		Username: m.Config.Username,
		Password: m.Config.Password,
		Opus:     true, // Enable Opus codec support.
	}); err != nil {
		return err
	}

	return nil
}
