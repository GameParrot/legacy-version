package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func DimensionData(io protocol.IO, pk *packet.DimensionData) {
	protocol.FuncIOSlice(io, &pk.Definitions, proto.MarshalDimensionDefinition)
}
