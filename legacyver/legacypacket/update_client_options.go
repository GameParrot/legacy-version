package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func UpdateClientOptions(io protocol.IO, pk *packet.UpdateClientOptions) {
	protocol.OptionalFunc(io, &pk.GraphicsMode, io.Uint8)
	if proto.IsProtoGTE(io, proto.ID975) {
		protocol.OptionalFunc(io, &pk.FilterProfanity, io.Bool)
	}
}
