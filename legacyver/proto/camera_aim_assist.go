package proto

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalCameraAimAssistCategory(r protocol.IO, x *protocol.CameraAimAssistCategory) {
	r.String(&x.Name)
	MarshalCameraAimAssistPriorities(r, &x.Priorities)
}

func MarshalCameraAimAssistPriorities(r protocol.IO, x *protocol.CameraAimAssistPriorities) {
	protocol.Slice(r, &x.Entities)
	protocol.Slice(r, &x.Blocks)
	if IsProtoGTE(r, ID898) {
		protocol.Slice(r, &x.BlockTags)
	}
	if IsProtoGTE(r, ID924) {
		protocol.Slice(r, &x.EntityTypeFamilies)
	}
	protocol.OptionalFunc(r, &x.EntityDefault, r.Int32)
	protocol.OptionalFunc(r, &x.BlockDefault, r.Int32)
}

func MarshalCameraAimAssistPreset(r protocol.IO, x *protocol.CameraAimAssistPreset) {
	r.String(&x.Identifier)
	protocol.FuncSlice(r, &x.BlockExclusions, r.String)
	if IsProtoGTE(r, ID898) {
		protocol.FuncSlice(r, &x.EntityExclusions, r.String)
		protocol.FuncSlice(r, &x.BlockTagExclusions, r.String)
	}
	if IsProtoGTE(r, ID924) {
		protocol.FuncSlice(r, &x.EntityTypeFamilyExclusions, r.String)
	}
	protocol.FuncSlice(r, &x.LiquidTargets, r.String)
	protocol.Slice(r, &x.ItemSettings)
	protocol.OptionalFunc(r, &x.DefaultItemSettings, r.String)
	protocol.OptionalFunc(r, &x.HandSettings, r.String)
}

func MarshalCameraRotationOption(r protocol.IO, x *protocol.CameraRotationOption) {
	r.Vec3(&x.Value)
	r.Float32(&x.Time)
	if IsProtoLT(r, ID924) {
		return
	}
	if IsReader(r) {
		x.EaseType = 0
	}
	if IsProtoLT(r, ID944) {
		marshalOptionalLegacyEaseType(r, x.EaseType, r.Uint8, func(value uint8) { x.EaseType = int32(value) })
	} else if IsProtoLT(r, ID975) {
		marshalOptionalLegacyEaseType(r, x.EaseType, r.String, func(value string) { easingTypeFromString(r, &x.EaseType, value) })
	} else {
		easingType := easingTypeToString(x.EaseType)
		r.String(&easingType)
		easingTypeFromString(r, &x.EaseType, easingType)
	}
}

func MarshalCameraSplineInstruction(r protocol.IO, x *protocol.CameraSplineInstruction) {
	r.Float32(&x.TotalTime)
	if IsProtoLT(r, ID924) {
		easeType, _ := x.SplineType.Value()
		r.Uint8(&easeType)
		x.SplineType = protocol.Option(easeType)
	} else {
		protocol.OptionalFunc(r, &x.SplineType, r.Uint8)
	}
	protocol.FuncSlice(r, &x.Curve, r.Vec3)
	if IsProtoLT(r, ID924) {
		keyFramesLegacy := make([]mgl32.Vec2, len(x.ProgressKeyFrames))
		for i, kf := range x.ProgressKeyFrames {
			keyFramesLegacy[i] = mgl32.Vec2{kf.Value, kf.Time}
		}
		protocol.FuncSlice(r, &keyFramesLegacy, r.Vec2)
		x.ProgressKeyFrames = make([]protocol.CameraProgressOption, len(keyFramesLegacy))
		for i, kf := range keyFramesLegacy {
			x.ProgressKeyFrames[i] = protocol.CameraProgressOption{
				Value: kf[0],
				Time:  kf[1],
			}
		}
	} else {
		protocol.FuncIOSlice(r, &x.ProgressKeyFrames, MarshalCameraProgressOption)
	}
	protocol.FuncIOSlice(r, &x.RotationOptions, MarshalCameraRotationOption)
	if IsProtoGTE(r, ID944) {
		protocol.OptionalFunc(r, &x.SplineIdentifier, r.String)
		protocol.OptionalFunc(r, &x.LoadFromJson, r.Bool)
	}
}

func MarshalCameraProgressOption(r protocol.IO, x *protocol.CameraProgressOption) {
	r.Float32(&x.Value)
	r.Float32(&x.Time)
	if IsReader(r) {
		x.EaseType = 0
	}
	if IsProtoLT(r, ID944) {
		marshalOptionalLegacyEaseType(r, x.EaseType, r.Uint8, func(value uint8) { x.EaseType = int32(value) })
	} else if IsProtoLT(r, ID975) {
		marshalOptionalLegacyEaseType(r, x.EaseType, r.String, func(value string) { easingTypeFromString(r, &x.EaseType, value) })
	} else {
		easingType := easingTypeToString(x.EaseType)
		r.String(&easingType)
		easingTypeFromString(r, &x.EaseType, easingType)
	}
}

func marshalOptionalLegacyEaseType[T comparable](r protocol.IO, easeType int32, marshal func(*T), decode func(T)) {
	var value protocol.Optional[T]
	if !IsReader(r) && easeType != 0 {
		var encoded any
		switch any(*new(T)).(type) {
		case uint8:
			encoded = uint8(easeType)
		case string:
			encoded = easingTypeToString(easeType)
		}
		value = protocol.Option(encoded.(T))
	}
	protocol.OptionalFunc(r, &value, marshal)
	if decoded, ok := value.Value(); ok {
		decode(decoded)
	}
}

func MarshalCameraSplineDefinition(r protocol.IO, x *protocol.CameraSplineDefinition) {
	r.String(&x.Name)
	if IsProtoGTE(r, ID944) {
		r.Float32(&x.TotalTime)
		protocol.OptionalFunc(r, &x.SplineType, r.String)
		protocol.FuncSlice(r, &x.ControlPoints, r.Vec3)
		protocol.FuncIOSlice(r, &x.ProgressKeyFrames, MarshalCameraProgressOption)
		protocol.FuncIOSlice(r, &x.RotationKeyFrames, MarshalCameraRotationOption)
	} else {
		d := &protocol.CameraSplineInstruction{
			TotalTime:         x.TotalTime,
			Curve:             x.ControlPoints,
			ProgressKeyFrames: x.ProgressKeyFrames,
			RotationOptions:   x.RotationKeyFrames,
		}
		MarshalCameraSplineInstruction(r, d)
		x.TotalTime = d.TotalTime
		x.ControlPoints = d.Curve
		x.ProgressKeyFrames = d.ProgressKeyFrames
		x.RotationKeyFrames = d.RotationOptions
	}
}
