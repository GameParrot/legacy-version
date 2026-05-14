package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalAbilityData(r protocol.IO, x *protocol.AbilityData) {
	r.Int64(&x.EntityUniqueID)
	r.Uint8(&x.PlayerPermissions)
	r.Uint8(&x.CommandPermissions)
	FuncIOSliceUint8Length(r, &x.Layers, MarshalAbilityLayer)
}

func MarshalAbilityLayer(r protocol.IO, x *protocol.AbilityLayer) {
	r.Uint16(&x.Type)
	r.Uint32(&x.Abilities)
	r.Uint32(&x.Values)
	r.Float32(&x.FlySpeed)
	if IsProtoGTE(r, ID776) {
		r.Float32(&x.VerticalFlySpeed)
	}
	r.Float32(&x.WalkSpeed)
}
