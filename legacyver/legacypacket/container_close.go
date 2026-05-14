package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ContainerClose(io protocol.IO, pk *packet.ContainerClose) {
	io.Uint8(&pk.WindowID)
	if proto.IsProtoGTE(io, proto.ID685) {
		io.Uint8(&pk.ContainerType)
	}
	io.Bool(&pk.ServerSide)
}
