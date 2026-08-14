package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func InventorySlot(io protocol.IO, pk *packet.InventorySlot) {
	io.Varuint32(&pk.WindowID)
	io.Varuint32(&pk.Slot)
	if proto.IsProtoGTE(io, proto.ID729) {
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.OptionalFunc(io, &pk.Container, func(x *protocol.FullContainerName) { proto.MarshalFullContainerName(io, x) })
		} else {
			container, _ := pk.Container.Value()
			protocol.Single(io, &container)
			pk.Container = protocol.Option(container)
		}
	}
	if proto.IsProtoGTE(io, proto.ID748) {
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.OptionalFunc(io, &pk.StorageItem, func(x *protocol.ItemInstance) { proto.ItemInstanceNew(io, x) })
		} else {
			item, _ := pk.StorageItem.Value()
			io.ItemInstance(&item)
			pk.StorageItem = protocol.Option(item)
		}
	} else {
		if proto.IsProtoGTE(io, proto.ID712) {
			dynamicContainerSize := uint32(0)
			io.Varuint32(&dynamicContainerSize)
		}
	}
	if proto.IsProtoGTE(io, proto.ID975) {
		proto.ItemInstanceNew(io, &pk.NewItem)
	} else {
		io.ItemInstance(&pk.NewItem)
	}
}
