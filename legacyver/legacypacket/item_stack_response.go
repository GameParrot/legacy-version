package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ItemStackResponse(io protocol.IO, pk *packet.ItemStackResponse) {
	protocol.FuncIOSlice(io, &pk.Responses, proto.MarshalItemStackResponse)
}
