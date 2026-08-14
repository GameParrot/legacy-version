package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CameraInstruction(io protocol.IO, pk *packet.CameraInstruction) {
	protocol.OptionalFunc(io, &pk.Set, func(x *protocol.CameraInstructionSet) { proto.MarshalCameraInstructionSet(io, x) })
	protocol.OptionalFunc(io, &pk.Clear, io.Bool)
	protocol.OptionalMarshaler(io, &pk.Fade)
	if proto.IsProtoGTE(io, proto.ID712) {
		protocol.OptionalMarshaler(io, &pk.Target)
		protocol.OptionalFunc(io, &pk.RemoveTarget, io.Bool)
	}
	if proto.IsProtoGTE(io, proto.ID827) {
		protocol.OptionalFunc(io, &pk.FieldOfView, func(x *protocol.CameraInstructionFieldOfView) { proto.MarshalCameraInstructionFieldOfView(io, x) })
	}
	if proto.IsProtoGTE(io, proto.ID859) {
		protocol.OptionalFunc(io, &pk.Spline, func(x *protocol.CameraSplineInstruction) { proto.MarshalCameraSplineInstruction(io, x) })
		protocol.OptionalFunc(io, &pk.AttachToEntity, io.Int64)
		protocol.OptionalFunc(io, &pk.DetachFromEntity, io.Bool)
	}
}
