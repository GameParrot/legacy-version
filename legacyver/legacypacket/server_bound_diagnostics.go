package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ServerBoundDiagnostics(io protocol.IO, pk *packet.ServerBoundDiagnostics) {
	io.Float32(&pk.AverageFramesPerSecond)
	io.Float32(&pk.AverageServerSimTickTime)
	io.Float32(&pk.AverageClientSimTickTime)
	io.Float32(&pk.AverageBeginFrameTime)
	io.Float32(&pk.AverageInputTime)
	io.Float32(&pk.AverageRenderTime)
	io.Float32(&pk.AverageEndFrameTime)
	io.Float32(&pk.AverageRemainderTimePercent)
	io.Float32(&pk.AverageUnaccountedTimePercent)
	if proto.IsProtoGTE(io, proto.ID924) {
		if proto.IsProtoLT(io, proto.ID944) {
			protocol.SliceUint32Length(io, &pk.MemoryCategoryValues)
		} else {
			protocol.Slice(io, &pk.MemoryCategoryValues)
		}
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.Slice(io, &pk.EntityDiagnostics)
			protocol.Slice(io, &pk.SystemDiagnostics)
			if proto.IsProtoGTE(io, proto.ID2168) {
				protocol.FuncIOSlice(io, &pk.SystemCategories, func(io protocol.IO, x *protocol.SystemCategory) {
					io.String(&x.CategoryName)
					io.Uint64(&x.SystemIndex)
				})
			}
			if proto.IsProtoGTE(io, proto.ID1001) {
				protocol.Slice(io, &pk.WhiskerScopes)
			}
		}
	}
}
