package legacypacket

import (
	_ "unsafe"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CommandOutput(io protocol.IO, pk *packet.CommandOutput) {
	proto.CommandOriginData(io, &pk.CommandOrigin)
	if proto.IsProtoGTE(io, proto.ID898) {
		outputTypeStr := commandOutputTypeToString(pk.OutputType)
		io.String(&outputTypeStr)
		commandOutputTypeFromString(io, &pk.OutputType, outputTypeStr)
		io.Uint32(&pk.SuccessCount)
	} else {
		io.Uint8(&pk.OutputType)
		io.Varuint32(&pk.SuccessCount)
	}
	protocol.FuncIOSlice(io, &pk.OutputMessages, proto.MarshalCommandOutputMessage)
	if proto.IsProtoGTE(io, proto.ID898) {
		protocol.OptionalFunc(io, &pk.DataSet, io.String)
	} else if pk.OutputType == packet.CommandOutputTypeDataSet {
		v, _ := pk.DataSet.Value()
		io.String(&v)
		if v != "" {
			pk.DataSet = protocol.Option(v)
		} else {
			pk.DataSet = protocol.Optional[string]{}
		}
	}
}

//go:linkname commandOutputTypeToString github.com/sandertv/gophertunnel/minecraft/protocol/packet.commandOutputTypeToString
func commandOutputTypeToString(x byte) string

//go:linkname commandOutputTypeFromString github.com/sandertv/gophertunnel/minecraft/protocol/packet.commandOutputTypeFromString
func commandOutputTypeFromString(io protocol.IO, x *byte, s string)
