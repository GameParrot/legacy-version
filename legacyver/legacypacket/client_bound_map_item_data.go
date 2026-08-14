package legacypacket

import (
	"image/color"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	mapUpdateFlagTexture = 1 << (iota + 1)
	mapUpdateFlagDecoration
	mapUpdateFlagInitialisation
)

func ClientBoundMapItemData(io protocol.IO, pk *packet.ClientBoundMapItemData) {
	io.Varint64(&pk.MapID)
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Uint8(&pk.Dimension)
		io.Bool(&pk.LockedMap)
		io.BlockPos(&pk.Origin)
		protocol.OptionalFunc(io, &pk.MapsIncludedIn, func(x *[]int64) { protocol.FuncSlice(io, x, io.Varint64) })
		protocol.OptionalFunc(io, &pk.Scale, io.Uint8)
		protocol.OptionalFunc(io, &pk.TrackedObjects, func(x *[]protocol.MapTrackedObject) {
			protocol.FuncIOSlice(io, x, proto.MarshalMapTrackedObject)
		})
		protocol.OptionalFunc(io, &pk.Decorations, func(x *[]protocol.MapDecoration) {
			protocol.FuncIOSlice(io, x, proto.MarshalMapDecoration)
		})
		protocol.OptionalFunc(io, &pk.Width, io.Varint32)
		protocol.OptionalFunc(io, &pk.Height, io.Varint32)
		protocol.OptionalFunc(io, &pk.XOffset, io.Varint32)
		protocol.OptionalFunc(io, &pk.YOffset, io.Varint32)
		protocol.OptionalFunc(io, &pk.Pixels, func(x *[]color.RGBA) { protocol.FuncSlice(io, x, io.BEARGB) })
		return
	}

	flags := uint32(0)
	if _, ok := pk.MapsIncludedIn.Value(); ok {
		flags |= mapUpdateFlagInitialisation
	}
	if _, ok := pk.TrackedObjects.Value(); ok {
		flags |= mapUpdateFlagDecoration
	}
	if _, ok := pk.Pixels.Value(); ok {
		flags |= mapUpdateFlagTexture
	}
	io.Varuint32(&flags)
	io.Uint8(&pk.Dimension)
	io.Bool(&pk.LockedMap)
	io.BlockPos(&pk.Origin)
	if flags&mapUpdateFlagInitialisation != 0 {
		x, _ := pk.MapsIncludedIn.Value()
		protocol.FuncSlice(io, &x, io.Varint64)
		pk.MapsIncludedIn = protocol.Option(x)
	}
	if flags&(mapUpdateFlagInitialisation|mapUpdateFlagDecoration|mapUpdateFlagTexture) != 0 {
		x, _ := pk.Scale.Value()
		io.Uint8(&x)
		pk.Scale = protocol.Option(x)
	}
	if flags&mapUpdateFlagDecoration != 0 {
		tracked, _ := pk.TrackedObjects.Value()
		protocol.FuncIOSlice(io, &tracked, proto.MarshalMapTrackedObject)
		pk.TrackedObjects = protocol.Option(tracked)
		decorations, _ := pk.Decorations.Value()
		protocol.FuncIOSlice(io, &decorations, proto.MarshalMapDecoration)
		pk.Decorations = protocol.Option(decorations)
	}
	if flags&mapUpdateFlagTexture != 0 {
		width, _ := pk.Width.Value()
		height, _ := pk.Height.Value()
		xOffset, _ := pk.XOffset.Value()
		yOffset, _ := pk.YOffset.Value()
		io.Varint32(&width)
		io.Varint32(&height)
		io.Varint32(&xOffset)
		io.Varint32(&yOffset)
		pk.Width, pk.Height = protocol.Option(width), protocol.Option(height)
		pk.XOffset, pk.YOffset = protocol.Option(xOffset), protocol.Option(yOffset)
		pixels, _ := pk.Pixels.Value()
		protocol.FuncSlice(io, &pixels, func(x *color.RGBA) {
			value := uint32(x.R) | uint32(x.G)<<8 | uint32(x.B)<<16 | uint32(x.A)<<24
			io.Varuint32(&value)
			*x = color.RGBA{R: byte(value), G: byte(value >> 8), B: byte(value >> 16), A: byte(value >> 24)}
		})
		pk.Pixels = protocol.Option(pixels)
	}
}
