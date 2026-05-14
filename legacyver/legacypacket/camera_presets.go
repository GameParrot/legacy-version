package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CameraPresets(io protocol.IO, pk *packet.CameraPresets) {
	protocol.FuncIOSlice(io, &pk.Presets, proto.MarshalCameraPreset)
}
