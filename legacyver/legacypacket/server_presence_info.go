package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ServerPresenceInfo(io protocol.IO, pk *packet.ServerPresenceInfo) {
	protocol.OptionalFunc(io, &pk.PresenceInfo, func(x *protocol.PresenceInfo) {
		proto.MarshalPresenceInfo(io, x)
	})
}
