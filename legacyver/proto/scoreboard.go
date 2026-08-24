package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalScoreboardEntry(r protocol.IO, x *protocol.ScoreboardEntry) {
	variant := uint32(x.IdentityType)
	r.Varuint32(&variant)
	x.IdentityType = byte(variant)
	typeNames := [...]string{"remove", "changeplayer", "changeentity", "changefakeplayer"}
	if variant >= uint32(len(typeNames)) {
		r.UnknownEnumOption(variant, "scoreboard entry variant")
		return
	}
	typeName := typeNames[variant]
	r.String(&typeName)
	r.Varint64(&x.EntryID)
	if x.IdentityType == protocol.ScoreboardIdentityRemove {
		objective := protocol.Optional[string]{}
		if x.ObjectiveName != "" {
			objective = protocol.Option(x.ObjectiveName)
		}
		if IsProtoGTE(r, ID2192) {
			protocol.OptionalFunc(r, &objective, r.String)
		} else {
			DoubleOptionalFunc(r, &objective, r.String)
		}
		x.ObjectiveName, _ = objective.Value()
		return
	}
	r.String(&x.ObjectiveName)
	r.Int32(&x.Score)
	switch x.IdentityType {
	case protocol.ScoreboardIdentityPlayer, protocol.ScoreboardIdentityEntity:
		r.ActorUniqueID(&x.EntityUniqueID)
	case protocol.ScoreboardIdentityFakePlayer:
		r.String(&x.DisplayName)
	}
}

func MarshalScoreboardIdentityEntry(r protocol.IO, x *protocol.ScoreboardIdentityEntry) {
	r.Varint64(&x.EntryID)
	protocol.OptionalFunc(r, &x.EntityUniqueID, r.ActorUniqueID)
}
