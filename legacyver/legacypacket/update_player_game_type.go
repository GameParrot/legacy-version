package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func UpdatePlayerGameType(io protocol.IO, pk *packet.UpdatePlayerGameType) {
	io.Varint32(&pk.GameType)
	io.Varint64(&pk.PlayerUniqueID)
	if proto.IsProtoGTE(io, proto.ID671) {
		io.Varuint64(&pk.Tick)
	}
}
