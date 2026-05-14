package proto

import (
	_ "unsafe"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalCameraPreset(r protocol.IO, x *protocol.CameraPreset) {
	r.String(&x.Name)
	r.String(&x.Parent)
	protocol.OptionalFunc(r, &x.PosX, r.Float32)
	protocol.OptionalFunc(r, &x.PosY, r.Float32)
	protocol.OptionalFunc(r, &x.PosZ, r.Float32)
	protocol.OptionalFunc(r, &x.RotX, r.Float32)
	protocol.OptionalFunc(r, &x.RotY, r.Float32)
	if IsProtoGTE(r, ID729) {
		protocol.OptionalFunc(r, &x.RotationSpeed, r.Float32)
		protocol.OptionalFunc(r, &x.SnapToTarget, r.Bool)
	}
	if IsProtoGTE(r, ID748) {
		protocol.OptionalFunc(r, &x.HorizontalRotationLimit, r.Vec2)
		protocol.OptionalFunc(r, &x.VerticalRotationLimit, r.Vec2)
		protocol.OptionalFunc(r, &x.ContinueTargeting, r.Bool)
	}
	if IsProtoGTE(r, ID766) {
		protocol.OptionalFunc(r, &x.TrackingRadius, r.Float32)
	}
	if IsProtoGTE(r, ID776) {
		protocol.OptionalFunc(r, &x.MinYawLimit, r.Float32)
		protocol.OptionalFunc(r, &x.MaxYawLimit, r.Float32)
	}
	protocol.OptionalFunc(r, &x.ViewOffset, r.Vec2)
	if IsProtoGTE(r, ID729) {
		protocol.OptionalFunc(r, &x.EntityOffset, r.Vec3)
	}
	protocol.OptionalFunc(r, &x.Radius, r.Float32)
	protocol.OptionalFunc(r, &x.AudioListener, r.Uint8)
	protocol.OptionalFunc(r, &x.PlayerEffects, r.Bool)
	if IsProtoGTE(r, ID748) && IsProtoLT(r, ID818) {
		alignTargetAndCameraForward := protocol.Optional[bool]{}
		protocol.OptionalFunc(r, &alignTargetAndCameraForward, r.Bool)
	}
	if IsProtoGTE(r, ID766) {
		protocol.OptionalMarshaler(r, &x.AimAssist)
	}
	if IsProtoGTE(r, ID800) {
		protocol.OptionalFunc(r, &x.ControlScheme, r.Uint8)
	}
}

func MarshalCameraInstructionSet(r protocol.IO, x *protocol.CameraInstructionSet) {
	r.Uint32(&x.Preset)
	protocol.OptionalMarshaler(r, &x.Ease)
	protocol.OptionalFunc(r, &x.Position, r.Vec3)
	protocol.OptionalFunc(r, &x.Rotation, r.Vec2)
	protocol.OptionalFunc(r, &x.Facing, r.Vec3)
	protocol.OptionalFunc(r, &x.ViewOffset, r.Vec2)
	protocol.OptionalFunc(r, &x.EntityOffset, r.Vec3)
	protocol.OptionalFunc(r, &x.Default, r.Bool)
	if IsProtoGTE(r, ID818) {
		r.Bool(&x.IgnoreStartingValuesComponent)
	}
}

func MarshalCameraInstructionFieldOfView(r protocol.IO, x *protocol.CameraInstructionFieldOfView) {
	r.Float32(&x.FieldOfView)
	r.Float32(&x.EaseTime)
	if IsProtoGTE(r, ID944) {
		easingType := easingTypeToString(x.EaseType)
		r.String(&easingType)
		easingTypeFromString(r, &x.EaseType, easingType)
	} else {
		easeType := uint8(x.EaseType)
		r.Uint8(&easeType)
		x.EaseType = int32(easeType)
	}
	r.Bool(&x.Clear)
}

//go:linkname easingTypeToString github.com/sandertv/gophertunnel/minecraft/protocol.easingTypeToString
func easingTypeToString(x int32) string

//go:linkname easingTypeFromString github.com/sandertv/gophertunnel/minecraft/protocol.easingTypeFromString
func easingTypeFromString(io protocol.IO, x *int32, s string)
