package legacypacket

import (
	_ "unsafe"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Animate(io protocol.IO, pk *packet.Animate) {
	if proto.IsProtoGTE(io, proto.ID898) {
		io.Uint8(&pk.ActionType)
	} else {
		v := int32(pk.ActionType)
		io.Varint32(&v)
		pk.ActionType = uint8(v)
	}
	io.Varuint64(&pk.EntityRuntimeID)
	if proto.IsProtoGTE(io, proto.ID859) || (pk.ActionType == 128 || pk.ActionType == 127) {
		io.Float32(&pk.Data)
	}
	if proto.IsProtoGTE(io, proto.ID898) {
		var swingSource protocol.Optional[string]
		if pk.SwingSource != 0 {
			swingSource = protocol.Option(swingSourceToString(pk.SwingSource))
		}
		protocol.OptionalFunc(io, &swingSource, io.String)
		if val, ok := swingSource.Value(); ok {
			swingSourceFromString(io, &pk.SwingSource, val)
		}
	}
}

//go:linkname swingSourceFromString github.com/sandertv/gophertunnel/minecraft/protocol/packet.swingSourceFromString
func swingSourceFromString(io protocol.IO, x *uint8, s string)

//go:linkname swingSourceToString github.com/sandertv/gophertunnel/minecraft/protocol/packet.swingSourceToString
func swingSourceToString(x uint8) string
