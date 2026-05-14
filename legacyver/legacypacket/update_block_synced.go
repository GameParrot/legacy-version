package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func UpdateBlockSynced(io protocol.IO, pk *packet.UpdateBlockSynced) {
	proto.IOUBlockPos(io, &pk.Position)
	io.Varuint32(&pk.NewBlockRuntimeID)
	io.Varuint32(&pk.Flags)
	io.Varuint32(&pk.Layer)
	io.Varuint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.TransitionType)
}
