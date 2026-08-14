package legacypacket

import (
	"math"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	legacySubChunkRequestModeLimitless = math.MaxUint32
	legacySubChunkRequestModeLimited   = math.MaxUint32 - 1
)

func LevelChunk(io protocol.IO, pk *packet.LevelChunk) {
	io.ChunkPos(&pk.Position)
	io.Varint32(&pk.Dimension)
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Varuint32(&pk.SubChunkCount)
		protocol.OptionalFunc(io, &pk.SubChunkLimit, io.Varint32)
	} else {
		subChunkCount := pk.SubChunkCount
		if limit, ok := pk.SubChunkLimit.Value(); ok {
			if limit < 0 {
				subChunkCount = legacySubChunkRequestModeLimitless
			} else {
				subChunkCount = legacySubChunkRequestModeLimited
			}
		}
		io.Varuint32(&subChunkCount)
		if proto.IsReader(io) {
			pk.SubChunkCount = subChunkCount
		}
		if subChunkCount == legacySubChunkRequestModeLimited {
			limit, _ := pk.SubChunkLimit.Value()
			highest := uint16(limit)
			io.Uint16(&highest)
			if proto.IsReader(io) {
				pk.SubChunkLimit = protocol.Option(int32(highest))
			}
		} else if subChunkCount == legacySubChunkRequestModeLimitless {
			if proto.IsReader(io) {
				pk.SubChunkLimit = protocol.Option(int32(-1))
			}
		}
	}
	io.Bool(&pk.CacheEnabled)
	if proto.IsProtoGTE(io, proto.ID2168) || pk.CacheEnabled {
		protocol.FuncSlice(io, &pk.BlobHashes, io.Uint64)
	}
	io.ByteSlice(&pk.RawPayload)
}
