package legacypacket

import (
	"fmt"
	"strings"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Text(io protocol.IO, pk *packet.Text) {
	if proto.IsProtoLT(io, proto.ID898) {
		io.Uint8(&pk.TextType)
	}
	io.Bool(&pk.NeedsTranslation)
	if proto.IsProtoGTE(io, proto.ID898) {
		var categoryType uint8
		switch pk.TextType {
		case packet.TextTypeRaw, packet.TextTypeTip, packet.TextTypeSystem, packet.TextTypeObjectWhisper, packet.TextTypeObjectAnnouncement, packet.TextTypeObject:
			categoryType = packet.TextCategoryMessageOnly
		case packet.TextTypeChat, packet.TextTypeWhisper, packet.TextTypeAnnouncement:
			categoryType = packet.TextCategoryAuthoredMessage
		default:
			categoryType = packet.TextCategoryMessageWithParameters
		}
		io.Uint8(&categoryType)
		// Protocols 898-923 include a quirky Mojang text-category string block.
		// 924+ removed this and only keeps the uint8 category and text type.
		if proto.IsProtoLT(io, proto.ID924) {
			for _, v := range textCategoryConstants(categoryType) {
				stringConst(io, v)
			}
		}
		io.Uint8(&pk.TextType)
	}
	switch pk.TextType {
	case packet.TextTypeChat, packet.TextTypeWhisper, packet.TextTypeAnnouncement:
		io.String(&pk.SourceName)
		io.String(&pk.Message)
	case packet.TextTypeRaw, packet.TextTypeTip, packet.TextTypeSystem, packet.TextTypeObject, packet.TextTypeObjectWhisper, packet.TextTypeObjectAnnouncement:
		io.String(&pk.Message)
	case packet.TextTypeTranslation, packet.TextTypePopup, packet.TextTypeJukeboxPopup:
		io.String(&pk.Message)
		protocol.FuncSlice(io, &pk.Parameters, io.String)
	}
	if proto.IsProtoGTE(io, proto.ID898) {
		if len(pk.Message) == 0 {
			io.InvalidValue(pk.Message, "message", "string cannot be empty")
		}
	}
	io.String(&pk.XUID)
	io.String(&pk.PlatformChatID)
	if proto.IsProtoGTE(io, proto.ID685) {
		if proto.IsProtoGTE(io, proto.ID898) {
			protocol.OptionalFunc(io, &pk.FilteredMessage, io.String)
		} else {
			v, _ := pk.FilteredMessage.Value()
			io.String(&v)
			if v != "" {
				pk.FilteredMessage = protocol.Option(v)
			} else {
				pk.FilteredMessage = protocol.Optional[string]{}
			}
		}
	}
}

func textCategoryConstants(categoryType uint8) []string {
	switch categoryType {
	case packet.TextCategoryMessageOnly:
		return []string{"raw", "tip", "systemMessage", "textObjectWhisper", "textObjectAnnouncement", "textObject"}
	case packet.TextCategoryAuthoredMessage:
		return []string{"chat", "whisper", "announcement"}
	default:
		return []string{"translate", "popup", "jukeboxPopup"}
	}
}

func stringConst(io protocol.IO, expected string) {
	if proto.IsReader(io) {
		var got string
		io.String(&got)
		if !strings.EqualFold(got, expected) {
			io.InvalidValue(got, "text category constant", fmt.Sprintf("expected %q", expected))
		}
		return
	}
	v := expected
	io.String(&v)
}
