package connection

import (
	"reflect"

	"github.com/bytecntrl/gumble/gumble/proto/tcp"
)

var allowedTypesTCP = map[reflect.Type]uint16{
	reflect.TypeOf(&tcp.Version{}):                0,
	reflect.TypeOf(&tcp.UDPTunnel{}):              1,
	reflect.TypeOf(&tcp.Authenticate{}):           2,
	reflect.TypeOf(&tcp.Ping{}):                   3,
	reflect.TypeOf(&tcp.Reject{}):                 4,
	reflect.TypeOf(&tcp.ServerSync{}):             5,
	reflect.TypeOf(&tcp.ChannelRemove{}):          6,
	reflect.TypeOf(&tcp.ChannelState{}):           7,
	reflect.TypeOf(&tcp.UserRemove{}):             8,
	reflect.TypeOf(&tcp.UserState{}):              9,
	reflect.TypeOf(&tcp.BanList{}):                10,
	reflect.TypeOf(&tcp.TextMessage{}):            11,
	reflect.TypeOf(&tcp.PermissionDenied{}):       12,
	reflect.TypeOf(&tcp.ACL{}):                    13,
	reflect.TypeOf(&tcp.QueryUsers{}):             14,
	reflect.TypeOf(&tcp.CryptSetup{}):             15,
	reflect.TypeOf(&tcp.ContextActionModify{}):    16,
	reflect.TypeOf(&tcp.ContextAction{}):          17,
	reflect.TypeOf(&tcp.UserList{}):               18,
	reflect.TypeOf(&tcp.VoiceTarget{}):            19,
	reflect.TypeOf(&tcp.PermissionQuery{}):        20,
	reflect.TypeOf(&tcp.CodecVersion{}):           21,
	reflect.TypeOf(&tcp.UserStats{}):              22,
	reflect.TypeOf(&tcp.RequestBlob{}):            23,
	reflect.TypeOf(&tcp.ServerConfig{}):           24,
	reflect.TypeOf(&tcp.SuggestConfig{}):          25,
	reflect.TypeOf(&tcp.PluginDataTransmission{}): 26,
}

func GetTCPMessageType(msg any) (uint16, bool) {
	val, ok := allowedTypesTCP[reflect.TypeOf(msg)]
	return val, ok
}
