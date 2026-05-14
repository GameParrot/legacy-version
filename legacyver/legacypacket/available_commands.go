package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func AvailableCommands(io protocol.IO, pk *packet.AvailableCommands) {
	protocol.FuncSlice(io, &pk.EnumValues, io.String)
	protocol.FuncSlice(io, &pk.ChainedSubcommandValues, io.String)
	protocol.FuncSlice(io, &pk.Suffixes, io.String)
	protocol.FuncIOSlice(io, &pk.Enums, proto.CommandEnumContext{EnumValues: pk.EnumValues}.Marshal)
	protocol.FuncIOSlice(io, &pk.ChainedSubcommands, proto.MarshalChainedSubcommand)
	protocol.FuncIOSlice(io, &pk.Commands, proto.MarshalCommand)
	protocol.Slice(io, &pk.DynamicEnums)
	protocol.Slice(io, &pk.Constraints)
}
