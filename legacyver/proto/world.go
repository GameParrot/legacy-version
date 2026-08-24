package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalDimensionDefinition(r protocol.IO, x *protocol.DimensionDefinition) {
	r.String(&x.Name)
	if IsProtoGTE(r, ID2192) {
		r.Varint32(&x.MinimumY)
		r.Varint32(&x.HeightRange)
	} else {
		maximumY, minimumY := x.MinimumY+x.HeightRange, x.MinimumY
		r.Varint32(&maximumY)
		r.Varint32(&minimumY)
		x.MinimumY = minimumY
		x.HeightRange = maximumY - minimumY
	}
	r.Varint32(&x.Generator)
	if IsProtoGTE(r, ID975) {
		r.Varint32(&x.DimensionType)
	}
	if IsProtoGTE(r, ID2168) {
		r.UUID(&x.PackID)
	}
	if IsProtoGTE(r, ID2192) {
		r.String(&x.DefaultBiome)
	}
}
