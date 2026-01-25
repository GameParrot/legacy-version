package legacypacket

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type DimensionData struct {
	Definitions []protocol.DimensionDefinition
}

func (*DimensionData) ID() uint32 {
	return packet.IDDimensionData
}

func (pk *DimensionData) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Definitions)
}
