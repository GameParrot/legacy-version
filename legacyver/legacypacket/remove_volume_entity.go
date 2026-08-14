package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func RemoveVolumeEntity(io protocol.IO, pk *packet.RemoveVolumeEntity) {
	if proto.IsProtoGTE(io, proto.ID859) {
		io.ActorRuntimeIDVaruint32(&pk.EntityRuntimeID)
	} else {
		legacyRuntimeID := uint64(pk.EntityRuntimeID)
		io.Uint64(&legacyRuntimeID)
		pk.EntityRuntimeID = uint32(legacyRuntimeID)
	}
	io.Varint32(&pk.Dimension)
}
