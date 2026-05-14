package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerAction(io protocol.IO, pk *packet.PlayerAction) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.ActionType)
	proto.IOUBlockPos(io, &pk.BlockPosition)
	proto.IOUBlockPos(io, &pk.ResultPosition)
	io.Varint32(&pk.BlockFace)
}
