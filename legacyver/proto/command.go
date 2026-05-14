package proto

import (
	"math"
	_ "unsafe"

	"github.com/akmalfairuz/legacy-version/internal/typeconf"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalCommand(r protocol.IO, c *protocol.Command) {
	r.String(&c.Name)
	r.String(&c.Description)
	r.Uint16(&c.Flags)

	if IsProtoGTE(r, ID898) {
		permLevel := commandPermissionToString(c.PermissionLevel)
		r.String(&permLevel)
		commandPermissionFromString(r, &c.PermissionLevel, permLevel)
	} else {
		r.Uint8(&c.PermissionLevel)
	}
	r.Uint32(&c.AliasesOffset)
	if IsProtoGTE(r, ID898) {
		protocol.FuncSlice(r, &c.ChainedSubcommandOffsets, r.Uint32)
	} else {
		offsets := typeconf.SliceIntToSliceInt[uint32, uint16](c.ChainedSubcommandOffsets)
		protocol.FuncSlice(r, &offsets, r.Uint16)
		c.ChainedSubcommandOffsets = typeconf.SliceIntToSliceInt[uint16, uint32](offsets)
	}
	protocol.Slice(r, &c.Overloads)
}

// CommandEnumContext holds context required for encoding command enums.
type CommandEnumContext struct {
	EnumValues []string
}

// Marshal encodes/decodes a CommandEnum.
func (ctx CommandEnumContext) Marshal(r protocol.IO, x *protocol.CommandEnum) {
	r.String(&x.Type)
	if IsProtoGTE(r, ID898) {
		protocol.FuncSlice(r, &x.ValueIndices, r.Uint32)
	} else {
		protocol.FuncIOSlice(r, &x.ValueIndices, ctx.enumOption)
	}
}

// enumOption writes/reads a command enum option as a byte/uint16/uint32,
// depending on the amount of enum values.
func (ctx CommandEnumContext) enumOption(r protocol.IO, opt *uint32) {
	n := len(ctx.EnumValues)
	switch {
	case n <= math.MaxUint8:
		val := byte(*opt)
		r.Uint8(&val)
		*opt = uint32(val)
	case n <= math.MaxUint16:
		val := uint16(*opt)
		r.Uint16(&val)
		*opt = uint32(val)
	default:
		r.Uint32(opt)
	}
}

func MarshalChainedSubcommand(r protocol.IO, x *protocol.ChainedSubcommand) {
	r.String(&x.Name)
	protocol.FuncIOSlice(r, &x.Values, MarshalChainedSubcommandValue)
}

func MarshalChainedSubcommandValue(r protocol.IO, x *protocol.ChainedSubcommandValue) {
	if IsProtoGTE(r, ID898) {
		r.Varuint32(&x.Index)
		r.Varuint32(&x.Value)
	} else {
		v := uint16(x.Index)
		r.Uint16(&v)
		x.Index = uint32(v)
		vv := uint16(x.Value)
		r.Uint16(&vv)
		x.Value = uint32(vv)
	}
}

// CommandOriginData reads/writes a CommandOrigin x using IO r.
func CommandOriginData(r protocol.IO, x *protocol.CommandOrigin) {
	if IsProtoGTE(r, ID898) {
		originStr := commandOriginToString(x.Origin)
		r.String(&originStr)
		commandOriginFromString(r, &x.Origin, originStr)
	} else {
		r.Varuint32(&x.Origin)
	}
	r.UUID(&x.UUID)
	r.String(&x.RequestID)
	if IsProtoGTE(r, ID898) {
		r.Int64(&x.PlayerUniqueID)
	} else if x.Origin == protocol.CommandOriginDevConsole || x.Origin == protocol.CommandOriginTest {
		r.Varint64(&x.PlayerUniqueID)
	}
}

func MarshalCommandOutputMessage(r protocol.IO, x *protocol.CommandOutputMessage) {
	if IsProtoLT(r, ID898) {
		r.Bool(&x.Success)
	}
	r.String(&x.Message)
	if IsProtoGTE(r, ID898) {
		r.Bool(&x.Success)
	}
	protocol.FuncSlice(r, &x.Parameters, r.String)
}

//go:linkname commandPermissionToString github.com/sandertv/gophertunnel/minecraft/protocol.commandPermissionToString
func commandPermissionToString(x byte) string

//go:linkname commandPermissionFromString github.com/sandertv/gophertunnel/minecraft/protocol.commandPermissionFromString
func commandPermissionFromString(r protocol.IO, x *byte, s string)

//go:linkname commandOriginToString github.com/sandertv/gophertunnel/minecraft/protocol.commandOriginToString
func commandOriginToString(x uint32) string

//go:linkname commandOriginFromString github.com/sandertv/gophertunnel/minecraft/protocol.commandOriginFromString
func commandOriginFromString(r protocol.IO, x *uint32, s string)
