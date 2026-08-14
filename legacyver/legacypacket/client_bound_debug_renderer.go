package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientBoundDebugRenderer(io protocol.IO, pk *packet.ClientBoundDebugRenderer) {
	if proto.IsProtoGTE(io, proto.ID898) {
		typeName := "cleardebugmarkers"
		if pk.Type == packet.ClientBoundDebugRendererAddCube {
			typeName = "adddebugmarkercube"
		}
		io.String(&typeName)
		switch typeName {
		case "cleardebugmarkers":
			pk.Type = packet.ClientBoundDebugRendererClear
		case "adddebugmarkercube":
			pk.Type = packet.ClientBoundDebugRendererAddCube
		default:
			io.UnknownEnumOption(typeName, "client bound debug renderer type")
			return
		}
	} else {
		legacyType := pk.Type + 1
		io.Uint32(&legacyType)
		if legacyType < 1 || legacyType > 2 {
			io.UnknownEnumOption(legacyType, "client bound debug renderer type")
			return
		}
		pk.Type = legacyType - 1
	}
	if pk.Type == packet.ClientBoundDebugRendererAddCube {
		io.String(&pk.Text)
		io.Vec3(&pk.Position)
		io.Float32(&pk.Red)
		io.Float32(&pk.Green)
		io.Float32(&pk.Blue)
		io.Float32(&pk.Alpha)
		io.Uint64(&pk.Duration)
	}
}
