package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientBoundMapItemData(io protocol.IO, pk *packet.ClientBoundMapItemData) {
	io.Varint64(&pk.MapID)
	io.Varuint32(&pk.UpdateFlags)
	io.Uint8(&pk.Dimension)
	io.Bool(&pk.LockedMap)
	io.BlockPos(&pk.Origin)

	if pk.UpdateFlags&packet.MapUpdateFlagInitialisation != 0 {
		protocol.FuncSlice(io, &pk.MapsIncludedIn, io.Varint64)
	}
	if pk.UpdateFlags&(packet.MapUpdateFlagInitialisation|packet.MapUpdateFlagDecoration|packet.MapUpdateFlagTexture) != 0 {
		io.Uint8(&pk.Scale)
	}
	if pk.UpdateFlags&packet.MapUpdateFlagDecoration != 0 {
		protocol.FuncIOSlice(io, &pk.TrackedObjects, proto.MarshalMapTrackedObject)
		protocol.Slice(io, &pk.Decorations)
	}
	if pk.UpdateFlags&packet.MapUpdateFlagTexture != 0 {
		io.Varint32(&pk.Width)
		io.Varint32(&pk.Height)
		io.Varint32(&pk.XOffset)
		io.Varint32(&pk.YOffset)
		protocol.FuncSlice(io, &pk.Pixels, io.VarRGBA)
	}
}
