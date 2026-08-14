package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PrimitiveShapes(io protocol.IO, pk *packet.PrimitiveShapes) {
	if proto.IsProto(io, proto.ID818) {
		count := uint32(len(pk.Shapes))
		io.Uint32(&count)
		protocol.FuncIOSliceOfLen(io, count, &pk.Shapes, proto.MarshalPrimitiveShape)
		return
	}
	protocol.FuncIOSlice(io, &pk.Shapes, proto.MarshalPrimitiveShape)
}
