package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerEnchantOptions(io protocol.IO, pk *packet.PlayerEnchantOptions) {
	protocol.FuncIOSlice(io, &pk.Options, proto.MarshalEnchantmentOption)
}
