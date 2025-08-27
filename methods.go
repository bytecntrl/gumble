package gumble

import (
	"google.golang.org/protobuf/proto"

	"github.com/bytecntrl/gumble/internal/proto/tcp"
)

type VersionOptions struct {
	VersionV1 uint32
	VersionV2 uint64
	Release   string
	Os        string
	OsVersion string
}

func (m *Mumble) Version(opts *VersionOptions) error {
	versionPacket := &tcp.Version{
		VersionV1: proto.Uint32(opts.VersionV1),
		VersionV2: proto.Uint64(opts.VersionV2),
		Release:   proto.String(opts.Release),
		Os:        proto.String(opts.Os),
		OsVersion: proto.String(opts.OsVersion),
	}

	if err := m.TCPConn.WritePacket(versionPacket); err != nil {
		return err
	}

	return nil
}

type UDPTunnelOptions struct {
	Packet []byte
}

func (m *Mumble) UDPTunnel(opts *UDPTunnelOptions) error {
	udpTunnelPacket := &tcp.UDPTunnel{
		Packet: opts.Packet,
	}

	if err := m.TCPConn.WritePacket(udpTunnelPacket); err != nil {
		return err
	}

	return nil
}

type AuthenticateOptions struct {
	Username     string
	Password     string
	Tokens       []string
	CeltVersions []int32
	Opus         bool
	ClientType   int32
}

func (m *Mumble) Authenticate(opts *AuthenticateOptions) error {
	authenticatePacket := &tcp.Authenticate{
		Username:     proto.String(opts.Username),
		Password:     proto.String(opts.Password),
		Tokens:       opts.Tokens,
		CeltVersions: opts.CeltVersions,
		Opus:         proto.Bool(opts.Opus),
		ClientType:   proto.Int32(opts.ClientType),
	}

	if err := m.TCPConn.WritePacket(authenticatePacket); err != nil {
		return err
	}

	return nil
}
