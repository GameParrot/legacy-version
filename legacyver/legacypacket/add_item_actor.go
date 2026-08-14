package legacypacket

import (
	legacyproto "github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func AddItemActor(r protocol.IO, pk *packet.AddItemActor) {
	if legacyproto.IsProtoGTE(r, legacyproto.ID2168) {
		r.ActorUniqueID(&pk.EntityUniqueID)
		r.ActorRuntimeID(&pk.EntityRuntimeID)
	} else {
		r.Varint64(&pk.EntityUniqueID)
		r.Varuint64(&pk.EntityRuntimeID)
	}
	r.ItemInstance(&pk.Item)
	r.Vec3(&pk.Position)
	r.Vec3(&pk.Velocity)
	legacyproto.MarshalEntityMetadata(r, &pk.EntityMetadata)
	r.Bool(&pk.FromFishing)
}
