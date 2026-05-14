package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func InventoryTransaction(io protocol.IO, pk *packet.InventoryTransaction) {
	io.Varint32(&pk.LegacyRequestID)
	if pk.LegacyRequestID != 0 {
		protocol.Slice(io, &pk.LegacySetItemSlots)
	}
	io.TransactionDataType(&pk.TransactionData)
	protocol.Slice(io, &pk.Actions)
	proto.MarshalInventoryTransactionData(io, pk.TransactionData)
}
