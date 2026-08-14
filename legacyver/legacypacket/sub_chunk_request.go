package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SubChunkRequest(io protocol.IO, pk *packet.SubChunkRequest) {
	io.Varint32(&pk.Dimension)
	if proto.IsProtoGTE(io, proto.ID1001) {
		protocol.FuncIOSlice(io, &pk.Offsets, proto.MarshalSubChunkOffset)
		if proto.IsProtoGTE(io, proto.ID2168) {
			io.SubChunkPos(&pk.Position)
		} else {
			io.Int32(&pk.Position[0])
			io.Int32(&pk.Position[1])
			io.Int32(&pk.Position[2])
		}
		return
	}
	proto.VarSubChunkPos(io, &pk.Position)
	proto.FuncIOSliceUint32Length(io, &pk.Offsets, proto.MarshalSubChunkOffset)
}
