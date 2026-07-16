package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func InventoryTransaction(io protocol.IO, pk *packet.InventoryTransaction) {
	io.Varint32(&pk.LegacyRequestID)
	if proto.IsProtoGTE(io, proto.ID1001) {
		hasLegacy := pk.LegacyRequestID < -1 && (pk.LegacyRequestID&1) == 0
		io.Bool(&hasLegacy)
		if hasLegacy {
			protocol.Slice(io, &pk.LegacySetItemSlots)
		}
		hasType := true
		io.Bool(&hasType)
		if !hasType {
			io.InvalidValue(hasType, "InventoryTransaction transaction type", "expected presence bool to be true")
		}
	} else {
		if pk.LegacyRequestID != 0 {
			protocol.Slice(io, &pk.LegacySetItemSlots)
		}
	}
	io.TransactionDataType(&pk.TransactionData)
	if proto.IsProtoGTE(io, proto.ID1001) {
		hasActions := true
		io.Bool(&hasActions)
		if !hasActions {
			io.InvalidValue(hasActions, "InventoryTransaction actions", "expected presence bool to be true")
		}
	}
	protocol.FuncIOSlice(io, &pk.Actions, proto.MarshalInventoryAction)
	proto.MarshalInventoryTransactionData(io, pk.TransactionData)
}
