package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CorrectPlayerMovePrediction(io protocol.IO, pk *packet.CorrectPlayerMovePrediction) {
	if proto.IsProtoGTE(io, proto.ID671) {
		io.Uint8(&pk.PredictionType)
	}
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Delta)
	if (proto.IsProtoGTE(io, proto.ID671) && proto.IsProtoLT(io, proto.ID827) && pk.PredictionType == packet.PredictionTypeVehicle) ||
		proto.IsProtoGTE(io, proto.ID827) {

		io.Vec2(&pk.Rotation)
		if proto.IsProtoGTE(io, proto.ID712) {
			protocol.OptionalFunc(io, &pk.VehicleAngularVelocity, io.Float32)
		}
	}
	io.Bool(&pk.OnGround)
	io.Varuint64(&pk.Tick)
}
