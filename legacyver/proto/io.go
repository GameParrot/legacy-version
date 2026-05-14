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
}

func NewReader(r *protocol.Reader, protocolID int32) *Reader {
	return &Reader{
		Reader:     r,
		protocolID: protocolID,
	}
}

func (r *Reader) SetProtocolID(protocolID int32) { r.protocolID = protocolID }
func (r *Reader) ProtocolID() int32              { return r.protocolID }

type Writer struct {
	*protocol.Writer

	protocolID int32
}

func NewWriter(w *protocol.Writer, protocolID int32) *Writer {
	return &Writer{w, protocolID}
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

func PlayerInventoryAction(io protocol.IO, x *protocol.UseItemTransactionData) {
	io.Varint32(&x.LegacyRequestID)
	if x.LegacyRequestID < -1 && (x.LegacyRequestID&1) == 0 {
		protocol.Slice(io, &x.LegacySetItemSlots)
	}
	protocol.Slice(io, &x.Actions)
	io.Varuint32(&x.ActionType)
	if IsProtoGTE(io, ID712) {
		io.Varuint32(&x.TriggerType)
	}
	io.BlockPos(&x.BlockPosition)
	io.Varint32(&x.BlockFace)
	io.Varint32(&x.HotBarSlot)
	io.ItemInstance(&x.HeldItem)
	io.Vec3(&x.Position)
	io.Vec3(&x.ClickedPosition)
	io.Varuint32(&x.BlockRuntimeID)
	if IsProtoGTE(io, ID712) {
		io.Varuint32(&x.ClientPrediction)
	}
}

func IOStackRequestAction(io protocol.IO, x *protocol.StackRequestAction) {
	if IsReader(io) {
		var id uint8
		io.Uint8(&id)
		if !lookupStackRequestAction(id, x) {
			io.UnknownEnumOption(id, "stack request action type")
			return
		}
	} else {
		var id byte
		if !lookupStackRequestActionType(*x, &id) {
			io.UnknownEnumOption(fmt.Sprintf("%T", *x), "stack request action type")
		}
		io.Uint8(&id)
	}
	MarshalStackRequestAction(io, *x)
}

func IORecipe(io protocol.IO, recipe *protocol.Recipe) {
	if IsReader(io) {
		var recipeType int32
		io.Varint32(&recipeType)
		if !lookupRecipe(recipeType, recipe) {
			io.UnknownEnumOption(recipeType, "crafting data recipe type")
			return
		}
		UnmarshalRecipe(io.(*protocol.Reader), *recipe)
	} else {
		var recipeType int32
		if !lookupRecipeType(*recipe, &recipeType) {
			io.UnknownEnumOption(fmt.Sprintf("%T", *recipe), "crafting recipe type")
		}
		io.Varint32(&recipeType)
		MarshalRecipe(io.(*protocol.Writer), *recipe)
	}
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
