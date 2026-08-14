package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerSkin(io protocol.IO, pk *packet.PlayerSkin) {
	io.UUID(&pk.UUID)
	proto.MarshalSkin(io, &pk.Skin)
	io.String(&pk.NewSkinName)
	io.String(&pk.OldSkinName)
	if proto.IsProtoLT(io, proto.ID2168) {
		io.Bool(&pk.Skin.Trusted)
	}
}
