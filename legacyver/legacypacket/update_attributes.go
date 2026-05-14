package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func UpdateAttributes(io protocol.IO, pk *packet.UpdateAttributes) {
	io.Varuint64(&pk.EntityRuntimeID)
	protocol.FuncIOSlice(io, &pk.Attributes, proto.MarshalAttribute)
	io.Varuint64(&pk.Tick)
}
