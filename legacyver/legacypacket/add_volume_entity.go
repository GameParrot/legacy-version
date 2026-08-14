package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func AddVolumeEntity(io protocol.IO, pk *packet.AddVolumeEntity) {
	if proto.IsProtoGTE(io, proto.ID859) {
		io.ActorRuntimeIDVaruint32(&pk.EntityRuntimeID)
	} else {
		legacyRuntimeID := uint64(pk.EntityRuntimeID)
		io.Uint64(&legacyRuntimeID)
		pk.EntityRuntimeID = uint32(legacyRuntimeID)
	}
	io.NBT(&pk.EntityMetadata, nbt.NetworkLittleEndian)
	io.String(&pk.EncodingIdentifier)
	io.String(&pk.InstanceIdentifier)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.Bounds[0])
		io.BlockPos(&pk.Bounds[1])
	} else {
		legacyUnsignedBlockPos(io, &pk.Bounds[0])
		legacyUnsignedBlockPos(io, &pk.Bounds[1])
	}
	io.Varint32(&pk.Dimension)
	io.String(&pk.EngineVersion)
}

func legacyUnsignedBlockPos(io protocol.IO, position *protocol.BlockPos) {
	io.Varint32(&position[0])
	y := uint32(position[1])
	io.Varuint32(&y)
	position[1] = int32(y)
	io.Varint32(&position[2])
}
