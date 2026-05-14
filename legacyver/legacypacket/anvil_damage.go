package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func AnvilDamage(io protocol.IO, pk *packet.AnvilDamage) {
	io.Uint8(&pk.Damage)
	proto.IOUBlockPos(io, &pk.AnvilPosition)
}
