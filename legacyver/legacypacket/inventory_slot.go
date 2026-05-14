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
			protocol.OptionalFuncIO(io, &pk.Container, proto.MarshalFullContainerName)
		} else {
			container, _ := pk.Container.Value()
			protocol.Single(io, &container)
			pk.Container = protocol.Option(container)
		}
	}
	if proto.IsProtoGTE(io, proto.ID748) {
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.OptionalFunc(io, &pk.StorageItem, io.ItemInstanceNew)
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
		io.ItemInstanceNew(&pk.NewItem)
	} else {
		io.ItemInstance(&pk.NewItem)
	}
}
