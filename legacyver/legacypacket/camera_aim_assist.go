package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CameraAimAssist(io protocol.IO, pk *packet.CameraAimAssist) {
	if proto.IsProtoGTE(io, proto.ID766) {
		io.String(&pk.Preset)
	}
	io.Vec2(&pk.Angle)
	io.Float32(&pk.Distance)
	io.Uint8(&pk.TargetMode)
	io.Uint8(&pk.Action)
	if proto.IsProtoGTE(io, proto.ID827) {
		io.Bool(&pk.ShowDebugRender)
	}
}
