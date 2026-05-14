package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ActorEvent(io protocol.IO, pk *packet.ActorEvent) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint8(&pk.EventType)
	io.Varint32(&pk.EventData)
	if proto.IsProtoGTE(io, proto.ID975) {
		protocol.OptionalFunc(io, &pk.FireAtPosition, io.Vec3)
	}
}
