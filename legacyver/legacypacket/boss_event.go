package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func BossEvent(io protocol.IO, pk *packet.BossEvent) {
	var playerUniqueID int64
	marshalBossEvent(io, pk, &playerUniqueID)
}

type TranslatedBossEvent struct {
	Pk             *packet.BossEvent
	PlayerUniqueID int64
}

func (pk *TranslatedBossEvent) ID() uint32 { return packet.IDBossEvent }

func (pk *TranslatedBossEvent) Marshal(io protocol.IO) {
	marshalBossEvent(io, pk.Pk, &pk.PlayerUniqueID)
}

func marshalBossEvent(io protocol.IO, pk *packet.BossEvent, playerUniqueID *int64) {
	io.Varint64(&pk.BossEntityUniqueID)
	if proto.IsProtoGTE(io, proto.ID1001) {
		if proto.IsProtoLT(io, proto.ID2192) {
			io.Varint64(playerUniqueID)
		}
		io.Uint8(&pk.EventType)
		io.String(&pk.BossBarTitle)
		io.String(&pk.FilteredBossBarTitle)
		io.Float32(&pk.HealthPercentage)
		io.Uint8(&pk.Colour)
		io.Uint8(&pk.Overlay)
		return
	}
	protocol.IntegerFunc(&pk.EventType, io.Varuint32)
	switch pk.EventType {
	case packet.BossEventShow:
		io.String(&pk.BossBarTitle)
		if proto.IsProtoGTE(io, proto.ID776) {
			io.String(&pk.FilteredBossBarTitle)
		}
		io.Float32(&pk.HealthPercentage)
		z := uint16(0)
		io.Uint16(&z)
		legacyBossEventColour(io, &pk.Colour)
		protocol.IntegerFunc(&pk.Overlay, io.Varuint32)
	case packet.BossEventRegisterPlayer, packet.BossEventUnregisterPlayer, packet.BossEventRequest:
		io.Varint64(playerUniqueID)
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
		z := uint16(0)
		io.Uint16(&z)
		legacyBossEventColour(io, &pk.Colour)
		protocol.IntegerFunc(&pk.Overlay, io.Varuint32)
	case packet.BossEventTexture:
		legacyBossEventColour(io, &pk.Colour)
		protocol.IntegerFunc(&pk.Overlay, io.Varuint32)
	default:
		io.UnknownEnumOption(pk.EventType, "boss event type")
	}
}

func legacyBossEventColour(io protocol.IO, colour *uint8) {
	if proto.IsReader(io) {
		var legacyColour uint32
		io.Varuint32(&legacyColour)
		switch legacyColour {
		case 0, 1, 2, 3, 4, 5:
			*colour = uint8(legacyColour)
		case 6:
			*colour = packet.BossEventColourWhite
		default:
			io.UnknownEnumOption(legacyColour, "legacy boss event colour")
		}
		return
	}
	var legacyColour uint32
	switch *colour {
	case packet.BossEventColourPink, packet.BossEventColourBlue, packet.BossEventColourRed,
		packet.BossEventColourGreen, packet.BossEventColourYellow, packet.BossEventColourPurple:
		legacyColour = uint32(*colour)
	case packet.BossEventColourWhite:
		legacyColour = 6
	default:
		io.UnknownEnumOption(*colour, "legacy boss event colour")
		return
	}
	io.Varuint32(&legacyColour)
}
