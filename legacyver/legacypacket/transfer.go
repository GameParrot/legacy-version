package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Transfer(io protocol.IO, pk *packet.Transfer) {
	io.String(&pk.Address)
	io.Uint16(&pk.Port)
	if proto.IsProtoGTE(io, proto.ID729) {
		io.Bool(&pk.ReloadWorld)
	}
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.OptionalFunc(io, &pk.GatheringJoinInfo, func(x *protocol.GatheringJoinInfo) { proto.MarshalGatheringJoinInfo(io, x) })
	}
}
