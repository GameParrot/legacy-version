package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CreativeContent(io protocol.IO, pk *packet.CreativeContent) {
	if proto.IsProtoGTE(io, proto.ID776) {
		protocol.Slice(io, &pk.Groups)
	}
	protocol.FuncIOSlice(io, &pk.Items, proto.MarshalCreativeItem)
}
