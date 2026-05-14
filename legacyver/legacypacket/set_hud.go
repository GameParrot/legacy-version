package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SetHud(io protocol.IO, pk *packet.SetHud) {
	if proto.IsProtoGTE(io, proto.ID786) {
		protocol.FuncSlice(io, &pk.Elements, io.Varint32)
		io.Varint32(&pk.Visibility)
	} else {
		elements := make([]uint8, len(pk.Elements))
		for i, v := range pk.Elements {
			elements[i] = uint8(v)
		}
		protocol.FuncSlice(io, &elements, io.Uint8)
		for i, v := range elements {
			pk.Elements[i] = int32(v)
		}

		visibility := uint8(pk.Visibility)
		io.Uint8(&visibility)
		pk.Visibility = int32(visibility)
	}
}
