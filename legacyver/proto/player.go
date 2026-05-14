package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalPlayerListEntry(r protocol.IO, x *protocol.PlayerListEntry) {
	r.UUID(&x.UUID)
	r.Varint64(&x.EntityUniqueID)
	r.String(&x.Username)
	r.String(&x.XUID)
	r.String(&x.PlatformChatID)
	r.Int32(&x.BuildPlatform)
	protocol.Single(r, &x.Skin)
	r.Bool(&x.Teacher)
	r.Bool(&x.Host)
	r.Bool(&x.SubClient)
	if IsProtoGTE(r, ID800) {
		r.ARGB(&x.PlayerColour)
	}
}

// PlayerMoveSettings reads/writes PlayerMovementSettings x to/from IO r.
func PlayerMoveSettings(r protocol.IO, x *protocol.PlayerMovementSettings) {
	if IsProtoLT(r, ID818) {
		var two int32 = 2
		r.Varint32(&two)
	}
	r.Varint32(&x.RewindHistorySize)
	r.Bool(&x.ServerAuthoritativeBlockBreaking)
}
