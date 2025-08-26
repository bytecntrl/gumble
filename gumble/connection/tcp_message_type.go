package connection

import (
	"google.golang.org/protobuf/proto"

	"github.com/bytecntrl/gumble/gumble/proto/tcp"
)

func GetTCPMessageType(msg proto.Message) (uint16, bool) {
	var val uint16
	ok := true

	switch msg.(type) {
	case *tcp.Version:
		val = 0
	case *tcp.UDPTunnel:
		val = 1
	case *tcp.Authenticate:
		val = 2
	case *tcp.Ping:
		val = 3
	case *tcp.Reject:
		val = 4
	case *tcp.ServerSync:
		val = 5
	case *tcp.ChannelRemove:
		val = 6
	case *tcp.ChannelState:
		val = 7
	case *tcp.UserRemove:
		val = 8
	case *tcp.UserState:
		val = 9
	case *tcp.BanList:
		val = 10
	case *tcp.TextMessage:
		val = 11
	case *tcp.PermissionDenied:
		val = 12
	case *tcp.ACL:
		val = 13
	case *tcp.QueryUsers:
		val = 14
	case *tcp.CryptSetup:
		val = 15
	case *tcp.ContextActionModify:
		val = 16
	case *tcp.ContextAction:
		val = 17
	case *tcp.UserList:
		val = 18
	case *tcp.VoiceTarget:
		val = 19
	case *tcp.PermissionQuery:
		val = 20
	case *tcp.CodecVersion:
		val = 21
	case *tcp.UserStats:
		val = 22
	case *tcp.RequestBlob:
		val = 23
	case *tcp.ServerConfig:
		val = 24
	case *tcp.SuggestConfig:
		val = 25
	case *tcp.PluginDataTransmission:
		val = 26
	default:
		ok = false
	}

	return val, ok
}
