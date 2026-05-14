package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalDimensionDefinition(r protocol.IO, x *protocol.DimensionDefinition) {
	r.String(&x.Name)
	r.Varint32(&x.Range[0])
	r.Varint32(&x.Range[1])
	r.Varint32(&x.Generator)
	if IsProtoGTE(r, ID975) {
		r.Varint32(&x.DimensionType)
	}
}
