package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Disconnect(io protocol.IO, pk *packet.Disconnect) {
	io.Varint32(&pk.Reason)
	io.Bool(&pk.HideDisconnectionScreen)
	if !pk.HideDisconnectionScreen {
		io.String(&pk.Message)
		if proto.IsProtoGTE(io, proto.ID712) {
			io.String(&pk.FilteredMessage)
		}
	}
}
