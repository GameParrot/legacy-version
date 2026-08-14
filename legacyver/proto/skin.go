package proto

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

var personaPieceTypes = [...]string{
	"unknown", "persona_skeleton", "persona_body", "persona_skin", "persona_bottom", "persona_feet",
	"persona_dress", "persona_top", "persona_high_pants", "persona_hands", "persona_outerwear",
	"persona_facial_hair", "persona_mouth", "persona_eyes", "persona_hair", "persona_hood", "persona_back",
	"persona_face_accessory", "persona_head", "persona_legs", "persona_left_leg", "persona_right_leg",
	"persona_arms", "persona_left_arm", "persona_right_arm", "persona_capes", "persona_classic_skin",
	"persona_emote", "unsupported",
}

func MarshalSkin(r protocol.IO, x *protocol.Skin) {
	r.String(&x.SkinID)
	r.String(&x.PlayFabID)
	r.ByteSlice(&x.SkinResourcePatch)
	r.Uint32(&x.SkinImageWidth)
	r.Uint32(&x.SkinImageHeight)
	r.ByteSlice(&x.SkinData)
	legacySkinAnimations(r, &x.Animations)
	r.Uint32(&x.CapeImageWidth)
	r.Uint32(&x.CapeImageHeight)
	r.ByteSlice(&x.CapeData)
	r.ByteSlice(&x.SkinGeometry)
	r.ByteSlice(&x.GeometryDataEngineVersion)
	r.ByteSlice(&x.AnimationData)
	r.String(&x.CapeID)
	r.String(&x.FullID)
	if IsProtoGTE(r, ID2168) {
		r.Uint8(&x.ArmSize)
		r.BEARGB(&x.SkinColour)
	} else {
		armSize := "slim"
		if x.ArmSize == protocol.ArmSizeWide {
			armSize = "wide"
		}
		r.String(&armSize)
		if strings.EqualFold(armSize, "wide") {
			x.ArmSize = protocol.ArmSizeWide
		} else {
			x.ArmSize = protocol.ArmSizeSlim
		}
		skinColour := fmt.Sprintf("#%02x%02x%02x", x.SkinColour.R, x.SkinColour.G, x.SkinColour.B)
		r.String(&skinColour)
		x.SkinColour = parseLegacyColour(skinColour, false)
	}
	legacyPersonaPieces(r, &x.PersonaPieces)
	legacyPersonaTints(r, &x.PieceTintColours)
	r.Bool(&x.PremiumSkin)
	r.Bool(&x.PersonaSkin)
	r.Bool(&x.PersonaCapeOnClassicSkin)
	r.Bool(&x.PrimaryUser)
	r.Bool(&x.OverrideAppearance)
	if IsProtoGTE(r, ID2168) {
		trusted := "false"
		if x.Trusted {
			trusted = "true"
		}
		r.String(&trusted)
		x.Trusted = strings.EqualFold(trusted, "true")
		r.String(&x.ProfileHash)
	}
}

func legacySkinAnimations(r protocol.IO, x *[]protocol.SkinAnimation) {
	marshal := func(r protocol.IO, x *protocol.SkinAnimation) {
		r.Uint32(&x.ImageWidth)
		r.Uint32(&x.ImageHeight)
		r.ByteSlice(&x.ImageData)
		if IsProtoGTE(r, ID2168) {
			r.Varuint32(&x.AnimationType)
		} else {
			r.Uint32(&x.AnimationType)
		}
		r.Float32(&x.FrameCount)
		if IsProtoGTE(r, ID2168) {
			r.Varuint32(&x.ExpressionType)
		} else {
			r.Uint32(&x.ExpressionType)
		}
	}
	if IsProtoGTE(r, ID2168) {
		protocol.FuncIOSlice(r, x, marshal)
	} else {
		count := uint32(len(*x))
		r.Uint32(&count)
		protocol.FuncIOSliceOfLen(r, count, x, marshal)
	}
}

func legacyPersonaPieces(r protocol.IO, x *[]protocol.PersonaPiece) {
	marshal := func(r protocol.IO, x *protocol.PersonaPiece) {
		r.String(&x.PieceID)
		if IsProtoGTE(r, ID2168) {
			r.Uint32(&x.PieceType)
			r.UUID(&x.PackID)
			r.Bool(&x.Default)
			r.String(&x.ProductID)
			return
		}
		pieceType := "unknown"
		if int(x.PieceType) < len(personaPieceTypes) {
			pieceType = personaPieceTypes[x.PieceType]
		}
		r.String(&pieceType)
		for i, name := range personaPieceTypes {
			if name == pieceType {
				x.PieceType = uint32(i)
				break
			}
		}
		packID := x.PackID.String()
		r.String(&packID)
		if parsed, err := uuid.Parse(packID); err == nil {
			x.PackID = parsed
		}
		r.Bool(&x.Default)
		r.String(&x.ProductID)
	}
	if IsProtoGTE(r, ID2168) {
		protocol.FuncIOSlice(r, x, marshal)
	} else {
		count := uint32(len(*x))
		r.Uint32(&count)
		protocol.FuncIOSliceOfLen(r, count, x, marshal)
	}
}

func legacyPersonaTints(r protocol.IO, x *[]protocol.PersonaPieceTintColour) {
	if IsProtoGTE(r, ID2168) {
		protocol.FuncIOSlice(r, x, func(r protocol.IO, x *protocol.PersonaPieceTintColour) {
			wireType := strings.TrimPrefix(x.PieceType, "persona_")
			if x.PieceType == "persona_hand" {
				wireType = "hands"
			}
			r.String(&wireType)
			if IsReader(r) {
				if wireType == "hands" {
					x.PieceType = "persona_hand"
				} else if wireType == "unsupported" {
					x.PieceType = wireType
				} else {
					x.PieceType = "persona_" + wireType
				}
			}
			for i := range x.Colours {
				r.BEARGB(&x.Colours[i])
			}
		})
		return
	}
	count := uint32(len(*x))
	r.Uint32(&count)
	protocol.FuncIOSliceOfLen(r, count, x, func(r protocol.IO, x *protocol.PersonaPieceTintColour) {
		r.String(&x.PieceType)
		colourCount := uint32(len(x.Colours))
		r.Uint32(&colourCount)
		for i := uint32(0); i < colourCount; i++ {
			colourString := "#0"
			if int(i) < len(x.Colours) {
				c := x.Colours[i]
				colourString = fmt.Sprintf("#%02x%02x%02x%02x", c.A, c.R, c.G, c.B)
			}
			r.String(&colourString)
			if i < 4 {
				x.Colours[i] = parseLegacyColour(colourString, true)
			}
		}
	})
}

func parseLegacyColour(value string, argb bool) color.RGBA {
	hex := strings.TrimPrefix(value, "#")
	if hex == "0" || hex == "" {
		return color.RGBA{}
	}
	parsed, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return color.RGBA{}
	}
	if argb && len(hex) > 6 {
		return color.RGBA{A: byte(parsed >> 24), R: byte(parsed >> 16), G: byte(parsed >> 8), B: byte(parsed)}
	}
	return color.RGBA{R: byte(parsed >> 16), G: byte(parsed >> 8), B: byte(parsed), A: 0xff}
}
