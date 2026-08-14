package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func AnvilDamage(io protocol.IO, pk *packet.AnvilDamage) {
	if proto.IsProtoLT(io, proto.ID2168) {
		var damage uint8
		io.Uint8(&damage)
	}
	proto.IOUBlockPos(io, &pk.AnvilPosition)
}
