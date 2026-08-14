package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func MobArmourEquipment(io protocol.IO, pk *packet.MobArmourEquipment) {
	io.Varuint64(&pk.EntityRuntimeID)
	if proto.IsProtoGTE(io, proto.ID1001) {
		proto.ItemInstanceNew(io, &pk.Helmet)
		proto.ItemInstanceNew(io, &pk.Chestplate)
		proto.ItemInstanceNew(io, &pk.Leggings)
		proto.ItemInstanceNew(io, &pk.Boots)
		proto.ItemInstanceNew(io, &pk.Body)
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
