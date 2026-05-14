package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientCheatAbility(io protocol.IO, pk *packet.ClientCheatAbility) {
	proto.MarshalAbilityData(io, &pk.AbilityData)
}
