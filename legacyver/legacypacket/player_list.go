package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerList(io protocol.IO, pk *packet.PlayerList) {
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.FuncIOSlice(io, &pk.Entries, proto.MarshalPlayerListEntry)
		return
	}
	action := uint8(protocol.PlayerListActionAdd)
	if len(pk.Entries) != 0 {
		action = pk.Entries[0].ActionType
	}
	io.Uint8(&action)
	switch action {
	case protocol.PlayerListActionAdd:
		protocol.FuncIOSlice(io, &pk.Entries, proto.MarshalPlayerListEntry)
	case protocol.PlayerListActionRemove:
		protocol.FuncIOSlice(io, &pk.Entries, func(io protocol.IO, x *protocol.PlayerListEntry) { io.UUID(&x.UUID) })
	default:
		io.UnknownEnumOption(action, "player list action type")
		return
	}
	for i := range pk.Entries {
		pk.Entries[i].ActionType = action
	}
	if action == protocol.PlayerListActionAdd {
		for i := range pk.Entries {
			io.Bool(&pk.Entries[i].Skin.Trusted)
		}
	}
}
