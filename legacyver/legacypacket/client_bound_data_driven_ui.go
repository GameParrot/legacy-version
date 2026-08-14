package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientBoundDataDrivenUIShowScreen(io protocol.IO, pk *packet.ClientBoundDataDrivenUIShowScreen) {
	io.String(&pk.ScreenID)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.Uint32(&pk.FormID)
		protocol.OptionalFunc(io, &pk.DataInstanceID, io.Uint32)
	}
}

func ClientBoundDataDrivenUICloseScreen(io protocol.IO, pk *packet.ClientBoundDataDrivenUICloseScreen) {
	if proto.IsProtoGTE(io, proto.ID944) {
		protocol.OptionalFunc(io, &pk.FormID, io.Uint32)
	}
}
