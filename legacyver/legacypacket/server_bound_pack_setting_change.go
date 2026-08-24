package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ServerBoundPackSettingChange(io protocol.IO, pk *packet.ServerBoundPackSettingChange) {
	io.UUID(&pk.PackID)
	proto.MarshalPackSetting(io, &pk.PackSetting)
}
