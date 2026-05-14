package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ResourcePackStack(io protocol.IO, pk *packet.ResourcePackStack) {
	io.Bool(&pk.TexturePackRequired)
	if proto.IsProtoLT(io, proto.ID898) {
		var emptyResourcePack []protocol.StackResourcePack
		protocol.Slice(io, &emptyResourcePack)
	}
	protocol.Slice(io, &pk.TexturePacks)
	io.String(&pk.BaseGameVersion)
	protocol.SliceUint32Length(io, &pk.Experiments)
	io.Bool(&pk.ExperimentsPreviouslyToggled)
	if proto.IsProtoGTE(io, proto.ID671) {
		io.Bool(&pk.IncludeEditorPacks)
	}
}
