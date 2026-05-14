package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func BlockEvent(io protocol.IO, pk *packet.BlockEvent) {
	proto.IOUBlockPos(io, &pk.Position)
	io.Varint32(&pk.EventType)
	io.Varint32(&pk.EventData)
}
