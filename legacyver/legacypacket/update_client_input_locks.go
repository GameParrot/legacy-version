package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func UpdateClientInputLocks(io protocol.IO, pk *packet.UpdateClientInputLocks) {
	io.Varuint32(&pk.Locks)
	if proto.IsProtoLT(io, proto.ID944) {
		// The position was removed in 1.26.10. It cannot be recovered from the
		// current packet, so encode/decode the legacy field independently.
		var position mgl32.Vec3
		io.Vec3(&position)
	}
}
