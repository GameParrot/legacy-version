package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func GraphicsOverrideParameter(io protocol.IO, pk *packet.GraphicsOverrideParameter) {
	protocol.Slice(io, &pk.Values)
	if proto.IsProtoGTE(io, proto.ID924) {
		if proto.IsProtoGTE(io, proto.ID944) {
			protocol.OptionalFunc(io, &pk.FloatValue, io.Float32)
			protocol.OptionalFunc(io, &pk.Vec3Value, io.Vec3)
		} else {
			floatVal, _ := pk.FloatValue.Value()
			io.Float32(&floatVal)
			pk.FloatValue = protocol.Option(floatVal)
			vec3Val, _ := pk.Vec3Value.Value()
			io.Vec3(&vec3Val)
			pk.Vec3Value = protocol.Option(vec3Val)
		}
	}
	io.String(&pk.BiomeIdentifier)
	io.Uint8(&pk.ParameterType)
	io.Bool(&pk.Reset)
}
