package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientMovementPredictionSync(io protocol.IO, pk *packet.ClientMovementPredictionSync) {
	io.Bitset(&pk.ActorFlags, proto.EntityDataFlagsLength(proto.FetchProtoID(io)))
	io.Float32(&pk.BoundingBoxScale)
	io.Float32(&pk.BoundingBoxWidth)
	io.Float32(&pk.BoundingBoxHeight)
	io.Float32(&pk.MovementSpeed)
	io.Float32(&pk.UnderwaterMovementSpeed)
	io.Float32(&pk.LavaMovementSpeed)
	io.Float32(&pk.JumpStrength)
	io.Float32(&pk.Health)
	io.Float32(&pk.Hunger)
	io.Varint64(&pk.EntityUniqueID)
	if proto.IsProtoGTE(io, proto.ID786) {
		io.Bool(&pk.Flying)
	}
}
