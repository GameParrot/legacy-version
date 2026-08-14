package legacypacket

import (
	legacyproto "github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ServerBoundDataStore(r protocol.IO, pk *packet.ServerBoundDataStore) {
	legacyproto.MarshalDataStoreUpdate(r, &pk.Update)
}

func ClientBoundDataStore(r protocol.IO, pk *packet.ClientBoundDataStore) {
	protocol.FuncIOSlice(r, &pk.Updates, legacyproto.MarshalDataStoreChangeEntry)
}
