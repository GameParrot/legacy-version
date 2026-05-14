package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func UpdateSubChunkBlocks(io protocol.IO, pk *packet.UpdateSubChunkBlocks) {
	proto.IOUBlockPos(io, &pk.Position)
	protocol.FuncIOSlice(io, &pk.Blocks, proto.MarshalBlockChangeEntry)
	protocol.FuncIOSlice(io, &pk.Extra, proto.MarshalBlockChangeEntry)
}
