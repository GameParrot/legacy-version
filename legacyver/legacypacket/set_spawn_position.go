package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SetSpawnPosition(io protocol.IO, pk *packet.SetSpawnPosition) {
	io.Varint32(&pk.SpawnType)
	proto.IOUBlockPos(io, &pk.Position)
	io.Varint32(&pk.Dimension)
	proto.IOUBlockPos(io, &pk.SpawnPosition)
}
