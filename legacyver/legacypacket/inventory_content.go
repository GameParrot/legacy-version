package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func InventoryContent(io protocol.IO, pk *packet.InventoryContent) {
	io.Varuint32(&pk.WindowID)
	protocol.FuncSlice(io, &pk.Content, io.ItemInstance)
	if proto.IsProtoGTE(io, proto.ID729) {
		proto.MarshalFullContainerName(io, &pk.Container)
	}
	if proto.IsProtoGTE(io, proto.ID748) {
		io.ItemInstance(&pk.StorageItem)
	} else {
		if proto.IsProtoGTE(io, proto.ID712) {
			dynamicContainerSize := uint32(0)
			io.Varuint32(&dynamicContainerSize)
		}
	}
}
