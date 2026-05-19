package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func MobArmourEquipment(io protocol.IO, pk *packet.MobArmourEquipment) {
	io.Varuint64(&pk.EntityRuntimeID)
	if proto.IsProtoGTE(io, proto.ID1001) {
		io.ItemInstanceNew(&pk.Helmet)
		io.ItemInstanceNew(&pk.Chestplate)
		io.ItemInstanceNew(&pk.Leggings)
		io.ItemInstanceNew(&pk.Boots)
		io.ItemInstanceNew(&pk.Body)
	} else {
		io.ItemInstance(&pk.Helmet)
		io.ItemInstance(&pk.Chestplate)
		io.ItemInstance(&pk.Leggings)
		io.ItemInstance(&pk.Boots)
		if proto.IsProtoGTE(io, proto.ID712) {
			io.ItemInstance(&pk.Body)
		}
	}
}
