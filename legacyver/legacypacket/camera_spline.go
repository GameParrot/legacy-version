package legacypacket

import (
	legacyproto "github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CameraSpline(r protocol.IO, pk *packet.CameraSpline) {
	protocol.FuncIOSlice(r, &pk.Splines, legacyproto.MarshalCameraSplineDefinition)
}
