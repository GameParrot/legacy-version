package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientCacheBlobStatus(io protocol.IO, pk *packet.ClientCacheBlobStatus) {
	if proto.IsProtoGTE(io, proto.ID1001) {
		protocol.FuncSlice(io, &pk.MissHashes, io.Uint64)
		protocol.FuncSlice(io, &pk.HitHashes, io.Uint64)
		return
	}
	missLen, hitLen := uint32(len(pk.MissHashes)), uint32(len(pk.HitHashes))
	io.Varuint32(&missLen)
	io.Varuint32(&hitLen)
	protocol.FuncSliceOfLen(io, missLen, &pk.MissHashes, io.Uint64)
	protocol.FuncSliceOfLen(io, hitLen, &pk.HitHashes, io.Uint64)
}
