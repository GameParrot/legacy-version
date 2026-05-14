package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/internal/typeconf"
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Event(io protocol.IO, pk *packet.Event) {
	io.Varint64(&pk.EntityRuntimeID)
	io.EventType(&pk.Event)
	if proto.IsProtoGTE(io, proto.ID898) {
		io.Bool(&pk.UsePlayerID)
	} else {
		v := typeconf.BoolToByte(pk.UsePlayerID)
		io.Uint8(&v)
		pk.UsePlayerID = typeconf.ByteToBool(v)
	}
	pk.Event.Marshal(io)
}
