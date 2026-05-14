package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func VoxelShapes(io protocol.IO, pk *packet.VoxelShapes) {
	protocol.Slice(io, &pk.Shapes)
	protocol.Slice(io, &pk.NameMap)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.Uint16(&pk.CustomShapeCount)
	}
}
