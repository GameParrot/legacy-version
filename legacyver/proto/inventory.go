package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalInventoryTransactionData(r protocol.IO, x protocol.InventoryTransactionData) {
	switch d := x.(type) {
	case *protocol.UseItemTransactionData:
		MarshalUseItemTransactionData(r, d)
	case *protocol.UseItemOnEntityTransactionData:
		MarshalUseItemOnEntityTransactionData(r, d)
	case *protocol.ReleaseItemTransactionData:
		MarshalReleaseItemTransactionData(r, d)

	default:
		x.Marshal(r)
	}
}

func MarshalUseItemTransactionData(r protocol.IO, x *protocol.UseItemTransactionData) {
	if IsProtoGTE(r, ID1001) {
		protocol.IntegerFunc(&x.ActionType, r.Varint32)
		protocol.IntegerFunc(&x.TriggerType, r.Uint8)
	} else {
		r.Varuint32(&x.ActionType)
		if IsProtoGTE(r, ID712) {
			r.Varuint32(&x.TriggerType)
		}
	}
	IOUBlockPos(r, &x.BlockPosition)
	if IsProtoGTE(r, ID1001) {
		protocol.IntegerFunc(&x.BlockFace, r.Uint8)
	} else {
		r.Varint32(&x.BlockFace)
	}
	r.Varint32(&x.HotBarSlot)
	if IsProtoGTE(r, ID1001) {
		r.ItemInstanceNew(&x.HeldItem)
	} else {
		r.ItemInstance(&x.HeldItem)
	}
	r.Vec3(&x.Position)
	r.Vec3(&x.ClickedPosition)
	r.Varuint32(&x.BlockRuntimeID)
	if IsProtoGTE(r, ID712) {
		r.Uint8(&x.ClientPrediction)
	}
	if IsProtoGTE(r, ID944) {
		r.Uint8(&x.ClientCooldownState)
	}
}

func MarshalReleaseItemTransactionData(r protocol.IO, data *protocol.ReleaseItemTransactionData) {
	if IsProtoGTE(r, ID1001) {
		r.Varint32(&data.ActionType)
	} else {
		protocol.IntegerFunc(&data.ActionType, r.Varuint32)
	}
	r.Varint32(&data.HotBarSlot)
	if IsProtoGTE(r, ID1001) {
		r.ItemInstanceNew(&data.HeldItem)
	} else {
		r.ItemInstance(&data.HeldItem)
	}
	r.Vec3(&data.HeadPosition)
}

func MarshalUseItemOnEntityTransactionData(r protocol.IO, data *protocol.UseItemOnEntityTransactionData) {
	r.Varuint64(&data.TargetEntityRuntimeID)
	if IsProtoGTE(r, ID1001) {
		r.Varint32(&data.ActionType)
	} else {
		protocol.IntegerFunc(&data.ActionType, r.Varuint32)
	}
	r.Varint32(&data.HotBarSlot)
	if IsProtoGTE(r, ID1001) {
		r.ItemInstanceNew(&data.HeldItem)
	} else {
		r.ItemInstance(&data.HeldItem)
	}
	r.Vec3(&data.Position)
	r.Vec3(&data.ClickedPosition)
}

func MarshalInventoryAction(r protocol.IO, x *protocol.InventoryAction) {
	r.Varuint32(&x.SourceType)
	if IsProtoGTE(r, ID1001) {
		present := true
		r.Bool(&present)
		if !present {
			r.InvalidValue(present, "InventoryAction container ID", "expected presence bool to be true")
		}
		hasContainerID := x.SourceType == protocol.InventoryActionSourceContainer || x.SourceType == protocol.InventoryActionSourceTODO
		r.Bool(&hasContainerID)
		if hasContainerID {
			r.Int8(&x.WindowID)
		}
		r.Bool(&present)
		if !present {
			r.InvalidValue(present, "InventoryAction source flags", "expected presence bool to be true")
		}
		hasFlags := x.SourceType == protocol.InventoryActionSourceWorld
		r.Bool(&hasFlags)
		if hasFlags {
			r.Varuint32(&x.SourceFlags)
		}
	} else {
		switch x.SourceType {
		case protocol.InventoryActionSourceContainer, protocol.InventoryActionSourceTODO:
			protocol.IntegerFunc(&x.WindowID, r.Varint32)
		case protocol.InventoryActionSourceWorld:
			r.Varuint32(&x.SourceFlags)
		}
	}
	r.Varuint32(&x.InventorySlot)
	if IsProtoGTE(r, ID1001) {
		r.ItemInstanceNew(&x.OldItem)
		r.ItemInstanceNew(&x.NewItem)
	} else {
		r.ItemInstance(&x.OldItem)
		r.ItemInstance(&x.NewItem)
	}
}
