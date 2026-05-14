package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CameraInstruction(io protocol.IO, pk *packet.CameraInstruction) {
	protocol.OptionalFuncIO(io, &pk.Set, proto.MarshalCameraInstructionSet)
	protocol.OptionalFunc(io, &pk.Clear, io.Bool)
	protocol.OptionalMarshaler(io, &pk.Fade)
	if proto.IsProtoGTE(io, proto.ID712) {
		protocol.OptionalMarshaler(io, &pk.Target)
		protocol.OptionalFunc(io, &pk.RemoveTarget, io.Bool)
	}
	if proto.IsProtoGTE(io, proto.ID827) {
		protocol.OptionalFuncIO(io, &pk.FieldOfView, proto.MarshalCameraInstructionFieldOfView)
	}
	if proto.IsProtoGTE(io, proto.ID859) {
		protocol.OptionalFuncIO(io, &pk.Spline, proto.MarshalCameraSplineInstruction)
		protocol.OptionalFunc(io, &pk.AttachToEntity, io.Int64)
		protocol.OptionalFunc(io, &pk.DetachFromEntity, io.Bool)
	}
}
