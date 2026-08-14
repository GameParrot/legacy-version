package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func GameRulesChanged(io protocol.IO, pk *packet.GameRulesChanged) {
	protocol.FuncIOSlice(io, &pk.GameRules, func(io protocol.IO, x *protocol.GameRule) {
		proto.MarshalGameRule(io, x, proto.IsProtoLT(io, proto.ID844))
	})
}
