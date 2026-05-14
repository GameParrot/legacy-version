package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CodeBuilderSource(io protocol.IO, pk *packet.CodeBuilderSource) {
	io.Uint8(&pk.Operation)
	io.Uint8(&pk.Category)
	if proto.IsProtoGTE(io, proto.ID685) {
		io.Uint8(&pk.CodeStatus)
	}
}
