package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func AgentAction(io protocol.IO, pk *packet.AgentAction) {
	io.String(&pk.Identifier)
	if proto.IsProtoGTE(io, proto.ID859) {
		io.Int32(&pk.Action)
	} else {
		io.Varint32(&pk.Action)
	}
	io.ByteSlice(&pk.Response)
}
