package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalInventoryTransactionData(r protocol.IO, x protocol.InventoryTransactionData) {
	switch d := x.(type) {
	case *protocol.UseItemTransactionData:
		MarshalUseItemTransactionData(r, d)
	default:
		x.Marshal(r)
	}
}

func MarshalUseItemTransactionData(r protocol.IO, x *protocol.UseItemTransactionData) {
	r.Varuint32(&x.ActionType)
	if IsProtoGTE(r, ID712) {
		r.Varuint32(&x.TriggerType)
	}
	IOUBlockPos(r, &x.BlockPosition)
	r.Varint32(&x.BlockFace)
	r.Varint32(&x.HotBarSlot)
	r.ItemInstance(&x.HeldItem)
	r.Vec3(&x.Position)
	r.Vec3(&x.ClickedPosition)
	r.Varuint32(&x.BlockRuntimeID)
	if IsProtoGTE(r, ID712) {
		r.Varuint32(&x.ClientPrediction)
	}
	if IsProtoGTE(r, ID944) {
		r.Uint8(&x.ClientCooldownState)
	}
}
