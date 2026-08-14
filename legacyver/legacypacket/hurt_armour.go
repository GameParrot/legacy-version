package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func HurtArmour(io protocol.IO, pk *packet.HurtArmour) {
	io.Varint32(&pk.Cause)
	io.Varint32(&pk.Damage)
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Varuint64(&pk.ArmourSlots)
	} else {
		legacySlots := int64(pk.ArmourSlots)
		io.Varint64(&legacySlots)
		pk.ArmourSlots = uint64(legacySlots)
	}
}
