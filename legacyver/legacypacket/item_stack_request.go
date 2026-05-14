package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ItemStackRequest(io protocol.IO, pk *packet.ItemStackRequest) {
	protocol.FuncIOSlice(io, &pk.Requests, proto.MarshalItemStackRequest)
}
