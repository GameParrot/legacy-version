package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func GameRulesChanged(io protocol.IO, pk *packet.GameRulesChanged) {
	if proto.IsProtoGTE(io, proto.ID844) {
		protocol.FuncSlice(io, &pk.GameRules, io.GameRule)
	} else {
		protocol.FuncSlice(io, &pk.GameRules, io.GameRuleLegacy)
	}
}
