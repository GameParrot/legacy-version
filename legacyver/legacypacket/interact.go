package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Interact(io protocol.IO, pk *packet.Interact) {
	io.Uint8(&pk.ActionType)
	io.Varuint64(&pk.TargetEntityRuntimeID)
	if proto.IsProtoGTE(io, proto.ID898) {
		protocol.OptionalFunc(io, &pk.Position, io.Vec3)
	} else if pk.ActionType == packet.InteractActionMouseOverEntity || pk.ActionType == packet.InteractActionLeaveVehicle {
		pos, _ := pk.Position.Value()
		io.Vec3(&pos)
		if pos != (mgl32.Vec3{}) {
			pk.Position = protocol.Option(pos)
		} else {
			pk.Position = protocol.Optional[mgl32.Vec3]{}
		}
	}
}
