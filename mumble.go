package gumble

import (
	"github.com/bytecntrl/gumble/gumble/model"
)

// Mumble is the struct used for the connection this the server
type Mumble struct {
	Config *model.Config // Config contains the client configurations
}

// NewMumble returns a new Mumble struct populated with the necessary defaults.
func NewMumble(config *model.Config) *Mumble {
	return &Mumble{Config: config}
}
