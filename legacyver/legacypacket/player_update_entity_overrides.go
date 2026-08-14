package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerUpdateEntityOverrides(io protocol.IO, pk *packet.PlayerUpdateEntityOverrides) {
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint32(&pk.PropertyIndex)
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Varuint32(&pk.Type)
		names := [...]string{"clearoverrides", "removeoverride", "setintoverride", "setfloatoverride"}
		if pk.Type >= uint32(len(names)) {
			io.UnknownEnumOption(pk.Type, "entity override type")
			return
		}
		name := names[pk.Type]
		io.String(&name)
	} else {
		legacyType := uint8(pk.Type)
		io.Uint8(&legacyType)
		pk.Type = uint32(legacyType)
	}
	switch pk.Type {
	case packet.PlayerUpdateEntityOverridesTypeClearAll, packet.PlayerUpdateEntityOverridesTypeRemove:
	case packet.PlayerUpdateEntityOverridesTypeInt:
		io.Int32(&pk.IntValue)
	case packet.PlayerUpdateEntityOverridesTypeFloat:
		io.Float32(&pk.FloatValue)
	default:
		io.UnknownEnumOption(pk.Type, "entity override type")
	}
}
