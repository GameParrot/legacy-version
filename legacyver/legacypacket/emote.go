package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Emote(io protocol.IO, pk *packet.Emote) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.String(&pk.EmoteID)
	if proto.IsProtoGTE(io, proto.ID729) {
		io.Varuint32(&pk.EmoteLength)
	}
	io.String(&pk.XUID)
	io.String(&pk.PlatformID)
	io.Uint8(&pk.Flags)
}
