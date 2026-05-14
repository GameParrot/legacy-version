package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func BlockActorData(io protocol.IO, pk *packet.BlockActorData) {
	proto.IOUBlockPos(io, &pk.Position)
	io.NBT(&pk.NBTData, nbt.NetworkLittleEndian)
}
