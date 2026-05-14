package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SetActorMotion(io protocol.IO, pk *packet.SetActorMotion) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Vec3(&pk.Velocity)
	if proto.IsProtoGTE(io, proto.ID662) {
		io.Varuint64(&pk.Tick)
	}
}
