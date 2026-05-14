package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func MobEffect(io protocol.IO, pk *packet.MobEffect) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint8(&pk.Operation)
	io.Varint32(&pk.EffectType)
	io.Varint32(&pk.Amplifier)
	io.Bool(&pk.Particles)
	io.Varint32(&pk.Duration)
	if proto.IsProtoGTE(io, proto.ID662) {
		if proto.IsProtoGTE(io, proto.ID748) {
			io.Varuint64(&pk.Tick)
		} else {
			io.Uint64(&pk.Tick)
		}

		if proto.IsProtoGTE(io, proto.ID898) {
			io.Bool(&pk.Ambient)
		}
	}
}
