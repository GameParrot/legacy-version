package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func SetScoreboardIdentity(io protocol.IO, pk *packet.SetScoreboardIdentity) {
	io.Uint8(&pk.ActionType)
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.FuncIOSlice(io, &pk.Entries, proto.MarshalScoreboardIdentityEntry)
		return
	}
	protocol.FuncIOSlice(io, &pk.Entries, func(io protocol.IO, x *protocol.ScoreboardIdentityEntry) {
		io.Varint64(&x.EntryID)
		if pk.ActionType == packet.ScoreboardIdentityActionRegister {
			entityID, _ := x.EntityUniqueID.Value()
			io.Varint64(&entityID)
			x.EntityUniqueID = protocol.Option(entityID)
		} else {
			x.EntityUniqueID = protocol.Optional[int64]{}
		}
	})
}
