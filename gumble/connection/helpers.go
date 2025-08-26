package connection

import (
	"net"
	"time"

	"github.com/bytecntrl/gumble/gumble"
)

func newDialer() *net.Dialer {
	return &net.Dialer{
		Timeout:   gumble.DialTimeout * time.Second,
		KeepAlive: gumble.DialKeepAlive * time.Second,
	}
}
