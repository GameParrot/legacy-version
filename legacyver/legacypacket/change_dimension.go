package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ChangeDimension(io protocol.IO, pk *packet.ChangeDimension) {
	io.Varint32(&pk.Dimension)
	io.Vec3(&pk.Position)
	io.Bool(&pk.Respawn)
	if proto.IsProtoGTE(io, proto.ID712) {
		protocol.OptionalFunc(io, &pk.LoadingScreenID, io.Uint32)
	}
}
