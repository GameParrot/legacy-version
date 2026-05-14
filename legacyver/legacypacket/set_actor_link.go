package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SetActorLink(io protocol.IO, pk *packet.SetActorLink) {
	proto.MarshalEntityLink(io, &pk.EntityLink)
}
