package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func BossEvent(io protocol.IO, pk *packet.BossEvent) {
	io.Varint64(&pk.BossEntityUniqueID)
	io.Varuint32(&pk.EventType)
	switch pk.EventType {
	case packet.BossEventShow:
		io.String(&pk.BossBarTitle)
		if proto.IsProtoGTE(io, proto.ID776) {
			io.String(&pk.FilteredBossBarTitle)
		}
		io.Float32(&pk.HealthPercentage)
		io.Uint16(&pk.ScreenDarkening)
		io.Varuint32(&pk.Colour)
		io.Varuint32(&pk.Overlay)
	case packet.BossEventRegisterPlayer, packet.BossEventUnregisterPlayer, packet.BossEventRequest:
		io.Varint64(&pk.PlayerUniqueID)
	case packet.BossEventHide:
		// No extra payload for this boss event type.
	case packet.BossEventHealthPercentage:
		io.Float32(&pk.HealthPercentage)
	case packet.BossEventTitle:
		io.String(&pk.BossBarTitle)
		if proto.IsProtoGTE(io, proto.ID776) {
			io.String(&pk.FilteredBossBarTitle)
		}
	case packet.BossEventAppearanceProperties:
		io.Uint16(&pk.ScreenDarkening)
		io.Varuint32(&pk.Colour)
		io.Varuint32(&pk.Overlay)
	case packet.BossEventTexture:
		io.Varuint32(&pk.Colour)
		io.Varuint32(&pk.Overlay)
	default:
		io.UnknownEnumOption(pk.EventType, "boss event type")
	}
}
