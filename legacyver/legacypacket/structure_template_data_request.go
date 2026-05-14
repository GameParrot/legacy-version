package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func StructureTemplateDataRequest(io protocol.IO, pk *packet.StructureTemplateDataRequest) {
	io.String(&pk.StructureName)
	proto.IOUBlockPos(io, &pk.Position)
	protocol.Single(io, &pk.Settings)
	io.Uint8(&pk.RequestType)
}
