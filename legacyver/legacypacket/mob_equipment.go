package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func MobEquipment(io protocol.IO, pk *packet.MobEquipment) {
	io.Varuint64(&pk.EntityRuntimeID)
	if proto.IsProtoGTE(io, proto.ID975) {
		io.ItemInstanceNew(&pk.NewItem)
	} else {
		io.ItemInstance(&pk.NewItem)
	}
	io.Uint8(&pk.InventorySlot)
	io.Uint8(&pk.HotBarSlot)
	io.Uint8(&pk.WindowID)
}
