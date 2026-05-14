package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CameraAimAssistPresets(io protocol.IO, pk *packet.CameraAimAssistPresets) {
	protocol.FuncIOSlice(io, &pk.Categories, proto.MarshalCameraAimAssistCategory)
	protocol.FuncIOSlice(io, &pk.Presets, proto.MarshalCameraAimAssistPreset)
	if proto.IsProtoGTE(io, proto.ID776) {
		io.Uint8(&pk.Operation)
	}
}
