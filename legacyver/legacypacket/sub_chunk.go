package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SubChunk(io protocol.IO, pk *packet.SubChunk) {
	io.Bool(&pk.CacheEnabled)
	io.Varint32(&pk.Dimension)
	io.SubChunkPos(&pk.Position)
	if pk.CacheEnabled {
		protocol.FuncIOSliceUint32Length(io, &pk.SubChunkEntries, proto.MarshalSubChunkEntry)
	} else {
		protocol.FuncIOSliceUint32Length(io, &pk.SubChunkEntries, proto.SubChunkEntryNoCache)
	}
}
