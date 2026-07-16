package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SubChunkRequest(io protocol.IO, pk *packet.SubChunkRequest) {
	io.Varint32(&pk.Dimension)
	if proto.IsProtoGTE(io, proto.ID1001) {
		protocol.Slice(io, &pk.Offsets)
		io.Int32(&pk.Position[0])
		io.Int32(&pk.Position[1])
		io.Int32(&pk.Position[2])
		return
	}
	io.SubChunkPos(&pk.Position)
	protocol.SliceUint32Length(io, &pk.Offsets)
}
