package legacypacket

import (
	"strings"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/samber/lo"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ItemRegistry(io protocol.IO, pk *packet.ItemRegistry) {
	if proto.IsProtoLT(io, proto.ID776) && proto.IsWriter(io) {
		items := lo.Filter(pk.Items, func(item protocol.ItemEntry, index int) bool {
			return !strings.HasPrefix(item.Name, "minecraft:") && item.ComponentBased
		}) // only custom items here
		protocol.FuncIOSlice(io, &items, proto.MarshalItemEntry)
		return
	}
	protocol.FuncIOSlice(io, &pk.Items, proto.MarshalItemEntry)
}
