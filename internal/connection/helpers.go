package connection

import (
	"net"
	"time"
)

const (
	DialTimeout   = 5
	DialKeepAlive = 30
)

func newDialer() *net.Dialer {
	return &net.Dialer{
		Timeout:   DialTimeout * time.Second,
		KeepAlive: DialKeepAlive * time.Second,
	}
}
