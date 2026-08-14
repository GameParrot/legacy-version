package proto

import (
	"image/color"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalPlayerListEntry(r protocol.IO, x *protocol.PlayerListEntry) {
	if IsProtoGTE(r, ID2168) {
		variant := uint32(0)
		if x.ActionType == protocol.PlayerListActionAdd {
			variant = 1
		}
		r.Varuint32(&variant)
		legacyAction := x.ActionType
		r.Uint8(&legacyAction)
		x.ActionType = protocol.PlayerListActionRemove
		if variant == 1 {
			x.ActionType = protocol.PlayerListActionAdd
		} else if variant != 0 {
			r.UnknownEnumOption(variant, "player list entry variant")
			return
		}
	}
	r.UUID(&x.UUID)
	if IsProtoGTE(r, ID2168) && x.ActionType == protocol.PlayerListActionRemove {
		return
	}
	r.Varint64(&x.EntityUniqueID)
	r.String(&x.Username)
	r.String(&x.XUID)
	r.String(&x.PlatformChatID)
	r.Int32(&x.BuildPlatform)
	MarshalSkin(r, &x.Skin)
	r.Bool(&x.Teacher)
	r.Bool(&x.Host)
	r.Bool(&x.SubClient)
	if IsProtoGTE(r, ID2168) {
		r.BEARGB(&x.PlayerColour)
	} else if IsProtoGTE(r, ID800) {
		legacyARGB(r, &x.PlayerColour)
	}
}

func legacyARGB(r protocol.IO, x *color.RGBA) {
	value := int32(x.A) | int32(x.R)<<8 | int32(x.G)<<16 | int32(x.B)<<24
	r.Int32(&value)
	*x = color.RGBA{A: byte(value), R: byte(value >> 8), G: byte(value >> 16), B: byte(value >> 24)}
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

func MarshalPlayerBlockAction(r protocol.IO, x *protocol.PlayerBlockAction) {
	r.Varint32(&x.Action)
	if IsProtoGTE(r, ID2168) || x.Action == protocol.PlayerActionStartBreak || x.Action == protocol.PlayerActionAbortBreak ||
		x.Action == protocol.PlayerActionCrackBreak || x.Action == protocol.PlayerActionPredictDestroyBlock || x.Action == protocol.PlayerActionContinueDestroyBlock {
		r.BlockPos(&x.BlockPos)
		r.Varint32(&x.Face)
	}
}
