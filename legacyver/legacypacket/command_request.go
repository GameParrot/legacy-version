package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/internal/typeconf"
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CommandRequest(io protocol.IO, pk *packet.CommandRequest) {
	io.String(&pk.CommandLine)
	proto.CommandOriginData(io, &pk.CommandOrigin)
	io.Bool(&pk.Internal)
	if proto.IsProtoGTE(io, proto.ID898) {
		io.String(&pk.Version)
	} else {
		v := typeconf.StringToInt[int32](pk.Version, 0)
		io.Varint32(&v)
		pk.Version = typeconf.IntToString[int32](v, "")
	}
}
