package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ChangeMobProperty(io protocol.IO, pk *packet.ChangeMobProperty) {
	if proto.IsProtoGTE(io, proto.ID859) {
		io.ActorUniqueID(&pk.EntityUniqueID)
	} else {
		legacyUniqueID := uint64(pk.EntityUniqueID)
		io.Uint64(&legacyUniqueID)
		pk.EntityUniqueID = int64(legacyUniqueID)
	}
	io.String(&pk.Property)
	io.Bool(&pk.BoolValue)
	io.String(&pk.StringValue)
	io.Varint32(&pk.IntValue)
	io.Float32(&pk.FloatValue)
}
