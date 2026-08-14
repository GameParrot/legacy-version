package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SubChunk(io protocol.IO, pk *packet.SubChunk) {
	io.Bool(&pk.CacheEnabled)
	io.Varint32(&pk.Dimension)
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.SubChunkPos(&pk.Position)
		protocol.FuncIOSlice(io, &pk.SubChunkEntries, proto.MarshalSubChunkEntry)
		return
	}
	proto.VarSubChunkPos(io, &pk.Position)
	if pk.CacheEnabled {
		proto.FuncIOSliceUint32Length(io, &pk.SubChunkEntries, proto.MarshalSubChunkEntry)
	} else {
		proto.FuncIOSliceUint32Length(io, &pk.SubChunkEntries, proto.SubChunkEntryNoCache)
	}
}
