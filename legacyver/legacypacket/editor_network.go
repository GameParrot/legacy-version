package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func EditorNetwork(io protocol.IO, pk *packet.EditorNetwork) {
	if proto.IsProtoGTE(io, proto.ID712) {
		io.Bool(&pk.RouteToManager)
	}
	io.NBT(&pk.Payload, nbt.NetworkLittleEndian)
}
