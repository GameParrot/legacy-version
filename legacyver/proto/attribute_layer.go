package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalEnvironmentAttributeData(r protocol.IO, x *protocol.EnvironmentAttributeData) {
	r.String(&x.AttributeName)
	protocol.OptionalMarshaler(r, &x.FromAttribute)
	protocol.Single(r, &x.Attribute)
	protocol.OptionalMarshaler(r, &x.ToAttribute)
	r.Uint32(&x.CurrentTransitionTicks)
	r.Uint32(&x.TotalTransitionTicks)
	if IsProtoLT(r, ID975) {
		r.Int32(&x.EaseType)
	} else {
		easingType := easingTypeToString(x.EaseType)
		r.String(&easingType)
		easingTypeFromString(r, &x.EaseType, easingType)
	}
	if IsProtoGTE(r, ID1001) {
		r.Uint32(&x.LocalTransitionTicks)
		r.Bool(&x.NoiseTransition)
	}
}

func MarshalAttributeLayerSettings(r protocol.IO, x *protocol.AttributeLayerSettings) {
	r.Int32(&x.Priority)
	if IsProtoLT(r, ID1001) {
		marshalLegacyAttributeLayerWeight(r, x, nil)
	} else {
		r.Float32(&x.FloatWeight)
	}
	r.Bool(&x.Enabled)
	r.Bool(&x.TransitionsPaused)
}

func MarshalAttributeLayerData(r protocol.IO, x *protocol.AttributeLayerData) {
	r.String(&x.Name)
	if IsProtoGTE(r, ID1001) {
		protocol.OptionalFunc(r, &x.NoiseName, r.String)
	}
	r.Varint32(&x.DimensionID)
	MarshalAttributeLayerSettings(r, &x.Settings)
	protocol.FuncIOSlice(r, &x.EnvironmentAttributes, MarshalEnvironmentAttributeData)
}

func marshalLegacyAttributeLayerWeight(r protocol.IO, settings *protocol.AttributeLayerSettings, noiseName *protocol.Optional[string]) {
	if IsReader(r) {
		var weightType uint32
		r.Varuint32(&weightType)
		switch weightType {
		case 0:
			r.Float32(&settings.FloatWeight)
		case 1:
			var name string
			r.String(&name)
			if noiseName != nil {
				*noiseName = protocol.Option(name)
			}
		default:
			r.UnknownEnumOption(weightType, "legacy attribute layer weight type")
		}
		return
	}
	if noiseName != nil {
		if name, ok := noiseName.Value(); ok {
			weightType := uint32(1)
			r.Varuint32(&weightType)
			r.String(&name)
			return
		}
	}
	weightType := uint32(0)
	r.Varuint32(&weightType)
	r.Float32(&settings.FloatWeight)
}
