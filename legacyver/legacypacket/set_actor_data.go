package legacypacket

import (
	legacyproto "github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SetActorData(r protocol.IO, pk *packet.SetActorData) {
	r.Varuint64(&pk.EntityRuntimeID)
	legacyproto.MarshalEntityMetadata(r, &pk.EntityMetadata)
	protocol.Single(r, &pk.EntityProperties)
	r.Varuint64(&pk.Tick)
}
