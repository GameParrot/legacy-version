package proto

import (
	"image/color"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalMapTrackedObject(r protocol.IO, x *protocol.MapTrackedObject) {
	r.Int32(&x.Type)
	if IsProtoGTE(r, ID2168) {
		protocol.OptionalFunc(r, &x.EntityUniqueID, r.ActorUniqueID)
		protocol.OptionalFunc(r, &x.BlockPosition, r.BlockPos)
		if x.Type != protocol.MapObjectTypeEntity && x.Type != protocol.MapObjectTypeBlock {
			r.UnknownEnumOption(x.Type, "map tracked object type")
		}
		return
	}
	switch x.Type {
	case protocol.MapObjectTypeEntity:
		entityID, _ := x.EntityUniqueID.Value()
		r.Varint64(&entityID)
		x.EntityUniqueID = protocol.Option(entityID)
	case protocol.MapObjectTypeBlock:
		position, _ := x.BlockPosition.Value()
		IOUBlockPos(r, &position)
		x.BlockPosition = protocol.Option(position)
	default:
		r.UnknownEnumOption(x.Type, "map tracked object type")
	}
}

func MarshalMapDecoration(r protocol.IO, x *protocol.MapDecoration) {
	r.Uint8(&x.Type)
	r.Uint8(&x.Rotation)
	r.Uint8(&x.X)
	r.Uint8(&x.Y)
	r.String(&x.Label)
	if IsProtoGTE(r, ID2168) {
		r.BEARGB(&x.Colour)
	} else {
		value := uint32(x.Colour.R) | uint32(x.Colour.G)<<8 | uint32(x.Colour.B)<<16 | uint32(x.Colour.A)<<24
		r.Varuint32(&value)
		x.Colour = color.RGBA{R: byte(value), G: byte(value >> 8), B: byte(value >> 16), A: byte(value >> 24)}
	}
}
