package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ShowStoreOffer(io protocol.IO, pk *packet.ShowStoreOffer) {
	if proto.IsProtoGTE(io, proto.ID859) {
		io.UUID(&pk.OfferID)
	} else {
		offerID := pk.OfferID.String()
		io.String(&offerID)
		if parsed, err := uuid.Parse(offerID); err == nil {
			pk.OfferID = parsed
		}
	}
	io.Uint8(&pk.Type)
}
