package proto

import (
	"fmt"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

type IO interface {
	protocol.IO

	SetProtocolID(protocolID int32)
	ProtocolID() int32
}

type Reader struct {
	*protocol.Reader

	protocolID int32
	shieldID   int32
	limits     bool
}

func NewReader(r *protocol.Reader, protocolID, shieldID int32, limits bool) *Reader {
	return &Reader{
		Reader:     r,
		protocolID: protocolID,
		shieldID:   shieldID,
		limits:     limits,
	}
}

func (r *Reader) SetProtocolID(protocolID int32) { r.protocolID = protocolID }
func (r *Reader) ProtocolID() int32              { return r.protocolID }

type Writer struct {
	*protocol.Writer

	protocolID int32
	shieldID   int32
}

func NewWriter(w *protocol.Writer, protocolID, shieldID int32) *Writer {
	return &Writer{Writer: w, protocolID: protocolID, shieldID: shieldID}
}

func (w *Writer) SetProtocolID(protocolID int32) { w.protocolID = protocolID }
func (w *Writer) ProtocolID() int32              { return w.protocolID }

func IsReader(r protocol.IO) bool {
	_, ok := r.(*Reader)
	return ok
}

func IsWriter(w protocol.IO) bool {
	_, ok := w.(*Writer)
	return ok
}

func EmptySlice[T any](io protocol.IO, slice *[]T) {
	if IsReader(io) {
		*slice = make([]T, 0)
	}
}

// DoubleOptionalFunc reads/writes the pre-1.26.50 optional value nested in an
// always-present outer optional. Gophertunnel removed this wire shape in 1.26.50.
func DoubleOptionalFunc[T any](r protocol.IO, x *protocol.Optional[T], f func(*T)) any {
	outer := true
	r.Bool(&outer)
	if outer {
		protocol.OptionalFunc(r, x, f)
	} else {
		*x = protocol.Optional[T]{}
	}
	return x
}

func PlayerInventoryAction(io protocol.IO, x *protocol.UseItemTransactionData) {
	io.Varint32(&x.LegacyRequestID)
	if IsProtoGTE(io, ID2168) {
		protocol.OptionalFunc(io, &x.LegacySetItemSlots, func(slots *[]protocol.LegacySetItemSlot) {
			protocol.FuncIOSlice(io, slots, marshalLegacySetItemSlot)
		})
	} else if x.LegacyRequestID < -1 && (x.LegacyRequestID&1) == 0 {
		items, _ := x.LegacySetItemSlots.Value()
		protocol.FuncIOSlice(io, &items, marshalLegacySetItemSlot)
		x.LegacySetItemSlots = protocol.Option(items)
	}
	if IsProtoGTE(io, ID2192) {
		protocol.FuncIOSlice(io, &x.Actions, MarshalInventoryAction)
	} else if IsProtoGTE(io, ID2168) {
		actions := protocol.Optional[[]protocol.InventoryAction]{}
		if x.Actions != nil {
			actions = protocol.Option(x.Actions)
		}
		DoubleOptionalFunc(io, &actions, func(actions *[]protocol.InventoryAction) {
			protocol.FuncIOSlice(io, actions, MarshalInventoryAction)
		})
		x.Actions, _ = actions.Value()
	} else {
		protocol.FuncIOSlice(io, &x.Actions, marshalLegacyPlayerInventoryAction)
	}
	if IsProtoGTE(io, ID2168) {
		protocol.IntegerFunc(&x.ActionType, io.Varint32)
		protocol.IntegerFunc(&x.TriggerType, io.Uint8)
	} else {
		io.Varuint32(&x.ActionType)
	}
	if IsProtoGTE(io, ID712) && IsProtoLT(io, ID2168) {
		io.Varuint32(&x.TriggerType)
	}
	IOUBlockPos(io, &x.BlockPosition)
	if IsProtoGTE(io, ID2168) {
		protocol.IntegerFunc(&x.BlockFace, io.Uint8)
	} else {
		io.Varint32(&x.BlockFace)
	}
	io.Varint32(&x.HotBarSlot)
	if IsProtoGTE(io, ID2192) {
		// The feature/26.50 Gophertunnel PlayerInventoryAction writer currently
		// omits this field even though UseItemTransactionData gained it. Treat
		// that omission as an upstream bug and keep the 26.50 wire field here.
		io.Uint8(&x.Hand)
	}
	io.ItemInstance(&x.HeldItem)
	io.Vec3(&x.Position)
	io.Vec3(&x.ClickedPosition)
	io.Varuint32(&x.BlockRuntimeID)
	if IsProtoGTE(io, ID712) {
		io.Uint8(&x.ClientPrediction)
	}
	if IsProtoGTE(io, ID944) {
		io.Uint8(&x.ClientCooldownState)
	}
}

func marshalLegacySetItemSlot(io protocol.IO, x *protocol.LegacySetItemSlot) {
	io.Uint8(&x.ContainerID)
	io.ByteSlice(&x.Slots)
}

func marshalLegacyPlayerInventoryAction(io protocol.IO, x *protocol.InventoryAction) {
	io.Varuint32(&x.SourceType)
	switch x.SourceType {
	case protocol.InventoryActionSourceContainer, protocol.InventoryActionSourceTODO:
		windowID, _ := x.WindowID.Value()
		protocol.IntegerFunc(&windowID, io.Varint32)
		x.WindowID = protocol.Option(windowID)
	case protocol.InventoryActionSourceWorld:
		flags, _ := x.SourceFlags.Value()
		io.Varuint32(&flags)
		x.SourceFlags = protocol.Option(flags)
	}
	io.Varuint32(&x.InventorySlot)
	io.ItemInstance(&x.OldItem)
	io.ItemInstance(&x.NewItem)
}

func IOStackRequestAction(io protocol.IO, x *protocol.StackRequestAction) {
	if IsReader(io) {
		var id uint8
		if IsProtoGTE(io, ID2168) {
			var variant uint32
			io.Varuint32(&variant)
			var legacyID uint8
			io.Uint8(&legacyID)
			id = uint8(variant)
			if variant >= uint32(protocol.StackRequestActionPlaceInContainer) {
				id += 2
			}
			if legacyID != id {
				io.InvalidValue(legacyID, "stack request action type", "does not match the variant it was sent under")
				return
			}
		} else {
			io.Uint8(&id)
		}
		if !lookupStackRequestAction(id, x) {
			io.UnknownEnumOption(id, "stack request action type")
			return
		}
	} else {
		var id byte
		if !lookupStackRequestActionType(*x, &id) {
			io.UnknownEnumOption(fmt.Sprintf("%T", *x), "stack request action type")
		}
		if IsProtoGTE(io, ID2168) {
			variant := uint32(id)
			if id > protocol.StackRequestActionTakeOutContainer {
				variant -= 2
			}
			io.Varuint32(&variant)
		}
		io.Uint8(&id)
	}
	MarshalStackRequestAction(io, *x)
}

func IOUBlockPos(io protocol.IO, x *protocol.BlockPos) {
	io.Varint32(&x[0])
	if IsProtoGTE(io, ID944) {
		io.Varint32(&x[1])
	} else {
		y := uint32(x[1])
		io.Varuint32(&y)
		x[1] = int32(y)
	}
	io.Varint32(&x[2])
}

func VarSubChunkPos(io protocol.IO, x *protocol.SubChunkPos) {
	io.Varint32(&x[0])
	io.Varint32(&x[1])
	io.Varint32(&x[2])
}

// FuncIOSliceUint8Length reads/writes a slice of T using a function with a uint8 length prefix.
func FuncIOSliceUint8Length[T any, S ~*[]T](r protocol.IO, x S, f func(protocol.IO, *T)) {
	count := uint8(len(*x))
	r.Uint8(&count)
	protocol.FuncIOSliceOfLen(r, uint32(count), x, f)
}

// FuncIOSliceUint16Length reads/writes a slice of T using a function with a uint16 length prefix.
func FuncIOSliceUint16Length[T any, S ~*[]T](r protocol.IO, x S, f func(protocol.IO, *T)) {
	count := uint16(len(*x))
	r.Uint16(&count)
	protocol.FuncIOSliceOfLen(r, uint32(count), x, f)
}

// FuncIOSliceUint32Length reads/writes a slice using a fixed uint32 length.
func FuncIOSliceUint32Length[T any, S ~*[]T](r protocol.IO, x S, f func(protocol.IO, *T)) {
	count := uint32(len(*x))
	r.Uint32(&count)
	protocol.FuncIOSliceOfLen(r, count, x, f)
}
