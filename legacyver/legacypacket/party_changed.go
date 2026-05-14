package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PartyChanged(io protocol.IO, pk *packet.PartyChanged) {
	if proto.IsProtoGTE(io, proto.ID975) {
		protocol.OptionalMarshaler(io, &pk.PartyInfo)
	} else {
		partyInfo, _ := pk.PartyInfo.Value()
		partyId := partyInfo.PartyID
		io.String(&partyId)
		partyInfo.PartyID = partyId
		pk.PartyInfo = protocol.Option(partyInfo)
	}
}
