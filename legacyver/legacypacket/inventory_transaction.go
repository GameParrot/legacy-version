package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func InventoryTransaction(io protocol.IO, pk *packet.InventoryTransaction) {
	io.Varint32(&pk.LegacyRequestID)
	present := pk.LegacyRequestID != 0
	if proto.IsProtoGTE(io, proto.ID998) {
		io.Bool(&present)
	}
	if present {
		protocol.Slice(io, &pk.LegacySetItemSlots)
	}
	if proto.IsProtoGTE(io, proto.ID998) {
		var b bool = true
		io.Bool(&b)
	}
	io.TransactionDataType(&pk.TransactionData)
	if proto.IsProtoGTE(io, proto.ID998) {
		var b bool = true
		io.Bool(&b)
		protocol.FuncIOSlice(io, &pk.Actions, proto.MarshalInventoryActionCereal)
		proto.MarshalInventoryTransactionDataCereal(io, pk.TransactionData)
	} else {
		protocol.Slice(io, &pk.Actions)
		proto.MarshalInventoryTransactionData(io, pk.TransactionData)
	}
}
