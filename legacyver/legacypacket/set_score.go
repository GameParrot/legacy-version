package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	legacyScoreModify = iota
	legacyScoreRemove
)

func SetScore(io protocol.IO, pk *packet.SetScore) {
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.FuncIOSlice(io, &pk.Entries, proto.MarshalScoreboardEntry)
		return
	}
	action := uint8(legacyScoreModify)
	if len(pk.Entries) != 0 && pk.Entries[0].IdentityType == protocol.ScoreboardIdentityRemove {
		action = legacyScoreRemove
	}
	io.Uint8(&action)
	if action == legacyScoreRemove {
		protocol.FuncIOSlice(io, &pk.Entries, func(io protocol.IO, x *protocol.ScoreboardEntry) {
			io.Varint64(&x.EntryID)
			io.String(&x.ObjectiveName)
			io.Int32(&x.Score)
			x.IdentityType = protocol.ScoreboardIdentityRemove
		})
		return
	}
	protocol.FuncIOSlice(io, &pk.Entries, func(io protocol.IO, x *protocol.ScoreboardEntry) {
		io.Varint64(&x.EntryID)
		io.String(&x.ObjectiveName)
		io.Int32(&x.Score)
		io.Uint8(&x.IdentityType)
		switch x.IdentityType {
		case protocol.ScoreboardIdentityPlayer, protocol.ScoreboardIdentityEntity:
			io.Varint64(&x.EntityUniqueID)
		case protocol.ScoreboardIdentityFakePlayer:
			io.String(&x.DisplayName)
		default:
			io.UnknownEnumOption(x.IdentityType, "scoreboard entry identity type")
		}
	})
}
