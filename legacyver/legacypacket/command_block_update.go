package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func CommandBlockUpdate(io protocol.IO, pk *packet.CommandBlockUpdate) {
	io.Bool(&pk.Block)
	if pk.Block {
		proto.IOUBlockPos(io, &pk.Position)
		io.Varuint32(&pk.Mode)
		io.Bool(&pk.NeedsRedstone)
		io.Bool(&pk.Conditional)
	} else {
		io.Varuint64(&pk.MinecartEntityRuntimeID)
	}
	io.String(&pk.Command)
	io.String(&pk.LastOutput)
	io.String(&pk.Name)
	if proto.IsProtoGTE(io, proto.ID776) {
		io.String(&pk.FilteredName)
	}
	io.Bool(&pk.ShouldTrackOutput)
	io.Uint32(&pk.TickDelay)
	io.Bool(&pk.ExecuteOnFirstTick)
}
