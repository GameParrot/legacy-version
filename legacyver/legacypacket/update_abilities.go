package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func UpdateAbilities(io protocol.IO, pk *packet.UpdateAbilities) {
	proto.MarshalAbilityData(io, &pk.AbilityData)
}
