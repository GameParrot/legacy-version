package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func JigsawStructureData(io protocol.IO, pk *packet.JigsawStructureData) {
	// Both layouts contain raw network-little-endian NBT. Before 1.26.40 the
	// packet exposed those bytes directly; the current packet exposes the map.
	if proto.IsProtoLT(io, proto.ID2168) {
		io.NBT(&pk.StructureData, nbt.NetworkLittleEndian)
		return
	}
	io.NBT(&pk.StructureData, nbt.NetworkLittleEndian)
}
