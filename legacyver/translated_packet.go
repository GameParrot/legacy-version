package legacyver

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type translatedPacket struct {
	pk        packet.Packet
	marshalFn func(io protocol.IO, pk packet.Packet)
}

// ID ...
func (pk *translatedPacket) ID() uint32 {
	return pk.pk.ID()
}

func (pk *translatedPacket) Marshal(io protocol.IO) {
	pk.marshalFn(io, pk.pk)
}
