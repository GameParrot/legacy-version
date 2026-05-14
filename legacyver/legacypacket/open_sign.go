package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func OpenSign(io protocol.IO, pk *packet.OpenSign) {
	proto.IOUBlockPos(io, &pk.Position)
	io.Bool(&pk.FrontSide)
}
