package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func AddVolumeEntity(io protocol.IO, pk *packet.AddVolumeEntity) {
	io.Varuint32(&pk.EntityRuntimeID)
	io.NBT(&pk.EntityMetadata, nbt.NetworkLittleEndian)
	io.String(&pk.EncodingIdentifier)
	io.String(&pk.InstanceIdentifier)
	proto.IOUBlockPos(io, &pk.Bounds[0])
	proto.IOUBlockPos(io, &pk.Bounds[1])
	io.Varint32(&pk.Dimension)
	io.String(&pk.EngineVersion)
}
