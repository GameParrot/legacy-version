package proto

import (
	"math"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Command holds the data that a command requires to be shown to a player client-side. The command is shown in
// the /help command and auto-completed using this data.
type Command struct {
	// Name is the name of the command. The command may be executed using this name, and will be shown in the
	// /help list with it. It currently seems that the client crashes if the Name contains uppercase letters.
	Name string
	// Description is the description of the command. It is shown in the /help list and when starting to write
	// a command.
	Description string
	// Flags is a combination of flags not currently known. Leaving the Flags field empty appears to work.
	Flags uint16
	// PermissionLevel is the command permission level that the player required to execute this command. The
	// field no longer seems to serve a purpose, as the client does not handle the execution of commands
	// anymore: The permissions should be checked server-side.
	PermissionLevel byte
	// AliasesOffset is the offset to a CommandEnum that holds the values that
	// should be used as aliases for this command.
	AliasesOffset uint32
	// ChainedSubcommandOffsets is a slice of offsets that all point to a different ChainedSubcommand from the
	// ChainedSubcommands slice in the AvailableCommands packet.
	ChainedSubcommandOffsets []uint16
	// Overloads is a list of command overloads that specify the ways in which a command may be executed. The
	// overloads may be completely different.
	Overloads []protocol.CommandOverload
}

func (c *Command) Marshal(r protocol.IO) {
	r.String(&c.Name)
	r.String(&c.Description)
	r.Uint16(&c.Flags)
	if IsProtoGTE(r, ID898) {
		s := "any"
		r.String(&s)
	} else {
		r.Uint8(&c.PermissionLevel)
	}
	r.Uint32(&c.AliasesOffset)
	if IsProtoGTE(r, ID898) {
		offsets32 := make([]uint32, len(c.ChainedSubcommandOffsets))
		for i, v := range c.ChainedSubcommandOffsets {
			offsets32[i] = uint32(v)
		}
		protocol.FuncSlice(r, &offsets32, r.Uint32)
		c.ChainedSubcommandOffsets = make([]uint16, len(offsets32))
		for i, v := range offsets32 {
			c.ChainedSubcommandOffsets[i] = uint16(v)
		}
	} else {
		protocol.FuncSlice(r, &c.ChainedSubcommandOffsets, r.Uint16)
	}
	protocol.Slice(r, &c.Overloads)
}

// CommandEnum represents an enum in a command usage. The enum typically has a type and a set of options that
// are valid. A value that is not one of the options results in a failure during execution.
type CommandEnum struct {
	// Type is the type of the command enum. The type will show up in the command usage as the type of the
	// argument if it has a certain amount of arguments, or when Options is set to true in the
	// command holding the enum.
	Type string
	// ValueIndices holds a list of indices that point to the EnumValues slice in the
	// AvailableCommandsPacket. These represent the options of the enum.
	ValueIndices []uint
}

// CommandEnumContext holds context required for encoding command enums.
type CommandEnumContext struct {
	EnumValues []string
}

// Marshal encodes/decodes a CommandEnum.
func (ctx CommandEnumContext) Marshal(r protocol.IO, x *CommandEnum) {
	r.String(&x.Type)
	protocol.FuncIOSlice(r, &x.ValueIndices, ctx.enumOption)
}

// enumOption writes/reads a command enum option as a byte/uint16/uint32,
// depending on the amount of enum values.
func (ctx CommandEnumContext) enumOption(r protocol.IO, opt *uint) {
	if IsProtoGTE(r, ID898) {
		val := uint32(*opt)
		r.Uint32(&val)
		*opt = uint(val)
		return
	}
	n := len(ctx.EnumValues)
	switch {
	case n <= math.MaxUint8:
		val := byte(*opt)
		r.Uint8(&val)
		*opt = uint(val)
	case n <= math.MaxUint16:
		val := uint16(*opt)
		r.Uint16(&val)
		*opt = uint(val)
	default:
		val := uint32(*opt)
		r.Uint32(&val)
		*opt = uint(val)
	}
}

// ChainedSubcommand represents a subcommand that can have chained commands, such as /execute which allows you to run
// another command as another entity or at a different position etc.
type ChainedSubcommand struct {
	// Name is the name of the chained subcommand and shows up in the list as a regular subcommand enum.
	Name string
	// Values contains the index and parameter type of the chained subcommand.
	Values []ChainedSubcommandValue
}

func (x *ChainedSubcommand) Marshal(r protocol.IO) {
	r.String(&x.Name)
	protocol.Slice(r, &x.Values)
}

// ChainedSubcommandValue represents the value for a chained subcommand argument.
type ChainedSubcommandValue struct {
	// Index is the index of the argument in the ChainedSubcommandValues slice from the AvailableCommands packet. This is
	// then used to set the type specified by the Value field below.
	Index uint16
	// Value is a combination of the flags above and specified the type of argument. Unlike regular parameter types,
	// this should NOT contain any of the special flags (valid, enum, suffixed or soft enum) but only the basic types.
	Value uint16
}

func (x *ChainedSubcommandValue) Marshal(r protocol.IO) {
	if IsProtoGTE(r, ID898) {
		index := uint32(x.Index)
		value := uint32(x.Value)
		r.Varuint32(&index)
		r.Varuint32(&value)
		x.Index = uint16(index)
		x.Value = uint16(value)
	} else {
		r.Uint16(&x.Index)
		r.Uint16(&x.Value)
	}
}

// CommandOrigin holds data that identifies the origin of the requesting of a command. It holds several
// fields that may be used to get specific information.
// When sent in a CommandRequest packet, the same CommandOrigin should be sent in a CommandOutput packet.
type CommandOrigin struct {
	// Origin is one of the values above that specifies the origin of the command. The origin may change,
	// depending on what part of the client actually called the command. The command may be issued by a
	// websocket server, for example.
	Origin uint32
	// UUID is a unique identifier for every instantiation of a command.
	UUID uuid.UUID
	// RequestID is an ID that identifies the request of the client. The server should send a CommandOrigin
	// with the same request ID to ensure it can be matched with the request by the caller of the command.
	// This is especially important for websocket servers and it seems that this field is only non-empty for
	// these websocket servers.
	RequestID string
	// PlayerUniqueID is an ID that identifies the player, the same as the one found in the AdventureSettings
	// packet. Filling it out with 0 seems to work.
	// PlayerUniqueID is only written if Origin is CommandOriginDevConsole or CommandOriginTest.
	PlayerUniqueID int64
}

// CommandOriginData reads/writes a CommandOrigin x using IO r.
func CommandOriginData(r protocol.IO, x *CommandOrigin) {
	if IsProtoGTE(r, ID898) {
		s := "player"
		r.String(&s)
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

// CommandOutputMessage represents a message sent by a command that holds the output of one of the commands
// executed.
type CommandOutputMessage struct {
	// Success indicates if the output message was one of a successful command execution. If set to true, the
	// output message is by default coloured white, whereas if set to false, the message is by default
	// coloured red.
	Success bool
	// Message is the message that is sent to the client in the chat window. It may either be simply a
	// message or a translated built-in string like 'commands.tp.success.coordinates', combined with specific
	// parameters below.
	Message string
	// Parameters is a list of parameters that serve to supply the message sent with additional information,
	// such as the position that a player was teleported to or the effect that was applied to an entity.
	// These parameters only apply for the Minecraft built-in command output.
	Parameters []string
}

// Marshal encodes/decodes a CommandOutputMessage.
func (x *CommandOutputMessage) Marshal(r protocol.IO) {
	r.String(&x.Message)
	r.Bool(&x.Success)
	protocol.FuncSlice(r, &x.Parameters, r.String)
}
