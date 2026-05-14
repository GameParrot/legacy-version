package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalMapTrackedObject(r protocol.IO, x *protocol.MapTrackedObject) {
	r.Int32(&x.Type)
	switch x.Type {
	case protocol.MapObjectTypeEntity:
		r.Varint64(&x.EntityUniqueID)
	case protocol.MapObjectTypeBlock:
		IOUBlockPos(r, &x.BlockPosition)
	default:
		r.UnknownEnumOption(x.Type, "map tracked object type")
	}
}
