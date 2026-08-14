package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerLocation(io protocol.IO, pk *packet.PlayerLocation) {
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Varint64(&pk.EntityUniqueID)
		protocol.IntegerFunc(&pk.Type, io.Varuint32)
		io.Varint32(&pk.Type)
	} else {
		io.Int32(&pk.Type)
		io.Varint64(&pk.EntityUniqueID)
	}
	if pk.Type == packet.PlayerLocationTypeCoordinates {
		io.Vec3(&pk.Position)
	} else if pk.Type != packet.PlayerLocationTypeHide {
		io.UnknownEnumOption(pk.Type, "player location type")
	}
}
