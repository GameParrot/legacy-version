package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func StopSound(io protocol.IO, pk *packet.StopSound) {
	io.String(&pk.SoundName)
	io.Bool(&pk.StopAll)
	if proto.IsProtoGTE(io, proto.ID712) {
		io.Bool(&pk.StopMusicLegacy)
	}
}
