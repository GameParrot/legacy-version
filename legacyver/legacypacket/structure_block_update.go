package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func StructureBlockUpdate(io protocol.IO, pk *packet.StructureBlockUpdate) {
	proto.IOUBlockPos(io, &pk.Position)
	io.String(&pk.StructureName)
	if proto.IsProtoGTE(io, proto.ID776) {
		io.String(&pk.FilteredStructureName)
	}
	io.String(&pk.DataField)
	io.Bool(&pk.IncludePlayers)
	io.Bool(&pk.ShowBoundingBox)
	io.Varint32(&pk.StructureBlockType)
	protocol.Single(io, &pk.Settings)
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Uint8(&pk.RedstoneSaveMode)
	} else {
		protocol.IntegerFunc(&pk.RedstoneSaveMode, io.Varint32)
	}
	io.Bool(&pk.ShouldTrigger)
	io.Bool(&pk.Waterlogged)
}
