package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func LecternUpdate(io protocol.IO, pk *packet.LecternUpdate) {
	io.Uint8(&pk.Page)
	io.Uint8(&pk.PageCount)
	proto.IOUBlockPos(io, &pk.Position)
}
