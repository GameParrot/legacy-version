package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ContainerOpen(io protocol.IO, pk *packet.ContainerOpen) {
	io.Uint8(&pk.WindowID)
	io.Uint8(&pk.ContainerType)
	proto.IOUBlockPos(io, &pk.ContainerPosition)
	io.Varint64(&pk.ContainerEntityUniqueID)
}
