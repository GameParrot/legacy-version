package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	TextTypeRaw = iota
	TextTypeChat
	TextTypeTranslation
	TextTypePopup
	TextTypeJukeboxPopup
	TextTypeTip
	TextTypeSystem
	TextTypeWhisper
	TextTypeAnnouncement
	TextTypeObjectWhisper
	TextTypeObject
	TextTypeObjectAnnouncement
)

// Text is sent by the client to the server to send chat messages, and by the server to the client to forward
// or send messages, which may be chat, popups, tips etc.
type Text struct {
	// TextType is the type of the text sent. When a client sends this to the server, it should always be
	// TextTypeChat. If the server sends it, it may be one of the other text types above.
	TextType byte
	// NeedsTranslation specifies if any of the messages need to be translated. It seems that where % is found
	// in translatable text types, these are translated regardless of this bool. Translatable text types
	// include TextTypeTranslation, TextTypeTip, TextTypePopup and TextTypeJukeboxPopup.
	NeedsTranslation bool
	// SourceName is the name of the source of the messages. This source is displayed in text types such as
	// the TextTypeChat and TextTypeWhisper, where typically the username is shown.
	SourceName string
	// Message is the message of the packet. This field is set for each TextType and is the main component of
	// the packet.
	Message string
	// Parameters is a list of parameters that should be filled into the message. These parameters are only
	// written if the type of the packet is TextTypeTranslation, TextTypeTip, TextTypePopup or TextTypeJukeboxPopup.
	Parameters []string
	// XUID is the XBOX Live user ID of the player that sent the message. It is only set for packets of
	// TextTypeChat. When sent to a player, the player will only be shown the chat message if a player with
	// this XUID is present in the player list and not muted, or if the XUID is empty.
	XUID string
	// PlatformChatID is an identifier only set for particular platforms when chatting (presumably only for
	// Nintendo Switch). It is otherwise an empty string, and is used to decide which players are able to
	// chat with each other.
	PlatformChatID string
	// FilteredMessage is a filtered version of Message with all the profanity removed. The client will use
	// this over Message if this field is not empty and they have the "Filter Profanity" setting enabled.
	FilteredMessage string
}

// ID ...
func (*Text) ID() uint32 {
	return packet.IDText
}

func (pk *Text) Marshal(io protocol.IO) {
	if proto.IsProtoGTE(io, proto.ID898) {
		io.Bool(&pk.NeedsTranslation)

		switch pk.TextType {
		case TextTypeChat, TextTypeWhisper, TextTypeAnnouncement:
			o := uint8(1)
			io.Uint8(&o)
			s := "chat"
			io.String(&s)
			s = "whisper"
			io.String(&s)
			s = "announcement"
			io.String(&s)
		case TextTypeRaw, TextTypeTip, TextTypeSystem, TextTypeObject, TextTypeObjectWhisper, TextTypeObjectAnnouncement:
			o := uint8(0)
			io.Uint8(&o)
			s := "raw"
			io.String(&s)
			s = "tip"
			io.String(&s)
			s = "systemMessage"
			io.String(&s)
			s = "textObjectWhisper"
			io.String(&s)
			s = "textObjectAnnouncement"
			io.String(&s)
			s = "textObject"
			io.String(&s)
		case TextTypeTranslation, TextTypePopup, TextTypeJukeboxPopup:
			o := uint8(2)
			io.Uint8(&o)
			s := "translate"
			io.String(&s)
			s = "popup"
			io.String(&s)
			s = "jukeboxPopup"
			io.String(&s)
		}

		io.Uint8(&pk.TextType)
	} else {
		io.Uint8(&pk.TextType)
		io.Bool(&pk.NeedsTranslation)
	}

	switch pk.TextType {
	case TextTypeChat, TextTypeWhisper, TextTypeAnnouncement:
		io.String(&pk.SourceName)
		io.String(&pk.Message)
	case TextTypeRaw, TextTypeTip, TextTypeSystem, TextTypeObject, TextTypeObjectWhisper, TextTypeObjectAnnouncement:
		io.String(&pk.Message)
	case TextTypeTranslation, TextTypePopup, TextTypeJukeboxPopup:
		io.String(&pk.Message)
		protocol.FuncSlice(io, &pk.Parameters, io.String)
	}

	io.String(&pk.XUID)
	io.String(&pk.PlatformChatID)

	if proto.IsProtoGTE(io, proto.ID898) {
		b := uint8(1)
		io.Uint8(&b)
		if b == 1 {
			io.String(&pk.FilteredMessage)
		}
	} else if proto.IsProtoGTE(io, proto.ID685) {
		io.String(&pk.FilteredMessage)
	}
}
