package proto

import (
	"fmt"

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
	case *protocol.NormalTransactionData, *protocol.MismatchTransactionData:
		// These variants carry no additional data.
	default:
		r.UnknownEnumOption(fmt.Sprintf("%T", x), "inventory transaction data type")
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
		ItemInstanceNew(r, &x.HeldItem)
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
		ItemInstanceNew(r, &data.HeldItem)
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
		ItemInstanceNew(r, &data.HeldItem)
	} else {
		r.ItemInstance(&data.HeldItem)
	}
	r.Vec3(&data.Position)
	r.Vec3(&data.ClickedPosition)
}

func MarshalInventoryAction(r protocol.IO, x *protocol.InventoryAction) {
	r.Varuint32(&x.SourceType)
	if IsProtoGTE(r, ID2168) {
		protocol.DoubleOptionalFunc(r, &x.WindowID, r.Int8)
		protocol.DoubleOptionalFunc(r, &x.SourceFlags, r.Varuint32)
	} else if IsProtoGTE(r, ID1001) {
		present := true
		r.Bool(&present)
		if !present {
			r.InvalidValue(present, "InventoryAction container ID", "expected presence bool to be true")
		}
		hasContainerID := x.SourceType == protocol.InventoryActionSourceContainer || x.SourceType == protocol.InventoryActionSourceTODO
		r.Bool(&hasContainerID)
		if hasContainerID {
			windowID, _ := x.WindowID.Value()
			r.Int8(&windowID)
			x.WindowID = protocol.Option(windowID)
		}
		r.Bool(&present)
		if !present {
			r.InvalidValue(present, "InventoryAction source flags", "expected presence bool to be true")
		}
		hasFlags := x.SourceType == protocol.InventoryActionSourceWorld
		r.Bool(&hasFlags)
		if hasFlags {
			flags, _ := x.SourceFlags.Value()
			r.Varuint32(&flags)
			x.SourceFlags = protocol.Option(flags)
		}
	} else {
		switch x.SourceType {
		case protocol.InventoryActionSourceContainer, protocol.InventoryActionSourceTODO:
			windowID, _ := x.WindowID.Value()
			protocol.IntegerFunc(&windowID, r.Varint32)
			x.WindowID = protocol.Option(windowID)
		case protocol.InventoryActionSourceWorld:
			flags, _ := x.SourceFlags.Value()
			r.Varuint32(&flags)
			x.SourceFlags = protocol.Option(flags)
		}
	}
	r.Varuint32(&x.InventorySlot)
	if IsProtoGTE(r, ID1001) && IsProtoLT(r, ID2168) {
		ItemInstanceNew(r, &x.OldItem)
		ItemInstanceNew(r, &x.NewItem)
	} else {
		r.ItemInstance(&x.OldItem)
		r.ItemInstance(&x.NewItem)
	}
}
