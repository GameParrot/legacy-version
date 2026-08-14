package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func MovePlayer(io protocol.IO, pk *packet.MovePlayer) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Vec3(&pk.Position)
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Float32(&pk.HeadYaw)
	io.Uint8(&pk.Mode)
	io.Bool(&pk.OnGround)
	io.Varuint64(&pk.RiddenEntityRuntimeID)
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.OptionalFunc(io, &pk.TeleportData, func(data *protocol.TeleportData) {
			io.Int32(&data.TeleportCause)
			io.Int32(&data.TeleportSourceEntityType)
		})
	} else if pk.Mode == packet.MoveModeTeleport {
		data, _ := pk.TeleportData.Value()
		io.Int32(&data.TeleportCause)
		io.Int32(&data.TeleportSourceEntityType)
		pk.TeleportData = protocol.Option(data)
	}
	io.Varuint64(&pk.Tick)
}
