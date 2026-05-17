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

func MarshalInventoryTransactionDataCereal(r protocol.IO, x protocol.InventoryTransactionData) {
	switch d := x.(type) {
	case *protocol.UseItemTransactionData:
		MarshalUseItemTransactionDataCereal(r, d)
	case *protocol.UseItemOnEntityTransactionData:
		MarshalUseItemOnEntityTransactionDataCereal(r, d)
	case *protocol.ReleaseItemTransactionData:
		MarshalReleaseItemTransactionDataCereal(r, d)
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

func MarshalUseItemTransactionDataCereal(r protocol.IO, x *protocol.UseItemTransactionData) {
	actionType := int32(x.ActionType)
	r.Varint32(&actionType)
	x.ActionType = uint32(actionType)
	triggerType := uint8(x.TriggerType)
	r.Uint8(&triggerType)
	x.TriggerType = uint32(triggerType)
	IOUBlockPos(r, &x.BlockPosition)
	face := int8(x.BlockFace)
	r.Int8(&face)
	x.BlockFace = int32(face)
	r.Varint32(&x.HotBarSlot)
	r.ItemInstanceNew(&x.HeldItem)
	r.Vec3(&x.Position)
	r.Vec3(&x.ClickedPosition)
	r.Varuint32(&x.BlockRuntimeID)
	clientPrediction := uint8(x.ClientPrediction)
	r.Uint8(&clientPrediction)
	x.ClientPrediction = uint32(clientPrediction)
	r.Uint8(&x.ClientCooldownState)
}

func MarshalReleaseItemTransactionDataCereal(r protocol.IO, data *protocol.ReleaseItemTransactionData) {
	r.Varuint32(&data.ActionType)
	r.Varint32(&data.HotBarSlot)
	r.ItemInstanceNew(&data.HeldItem)
	r.Vec3(&data.HeadPosition)
}

func MarshalUseItemOnEntityTransactionDataCereal(r protocol.IO, data *protocol.UseItemOnEntityTransactionData) {
	r.Varuint64(&data.TargetEntityRuntimeID)
	actionType := int32(data.ActionType)
	r.Varint32(&actionType)
	data.ActionType = uint32(actionType)
	r.Varint32(&data.HotBarSlot)
	r.ItemInstanceNew(&data.HeldItem)
	r.Vec3(&data.Position)
	r.Vec3(&data.ClickedPosition)
}

func MarshalInventoryActionCereal(r protocol.IO, x *protocol.InventoryAction) {
	r.Varuint32(&x.SourceType)
	a := false
	r.Bool(&a)
	bb := true
	r.Bool(&bb)
	if bb {
		windowId := int8(x.WindowID)
		r.Int8(&windowId)
		x.WindowID = int32(windowId)
	} else {
		var a bool
		r.Bool(&a)
	}
	sourceFlags := protocol.Option(x.SourceFlags)
	protocol.OptionalFunc(r, &sourceFlags, r.Varuint32)
	x.SourceFlags, _ = sourceFlags.Value()
	r.Varuint32(&x.InventorySlot)
	r.ItemInstanceNew(&x.OldItem)
	r.ItemInstanceNew(&x.NewItem)
}
