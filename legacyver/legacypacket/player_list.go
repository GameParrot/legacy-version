package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerList(io protocol.IO, pk *packet.PlayerList) {
	io.Uint8(&pk.ActionType)
	switch pk.ActionType {
	case packet.PlayerListActionAdd:
		protocol.FuncIOSlice(io, &pk.Entries, proto.MarshalPlayerListEntry)
	case packet.PlayerListActionRemove:
		protocol.FuncIOSlice(io, &pk.Entries, protocol.PlayerListRemoveEntry)
	default:
		io.UnknownEnumOption(pk.ActionType, "player list action type")
	}
	if pk.ActionType == packet.PlayerListActionAdd {
		for i := 0; i < len(pk.Entries); i++ {
			io.Bool(&pk.Entries[i].Skin.Trusted)
		}
	}
}
