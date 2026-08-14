package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ResourcePackClientResponse(io protocol.IO, pk *packet.ResourcePackClientResponse) {
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Varuint32(&pk.Response)
		names := [...]string{"cancel", "downloading", "downloadingfinished", "resourcepackstackfinished"}
		if pk.Response >= uint32(len(names)) {
			io.UnknownEnumOption(pk.Response, "resource pack response")
			return
		}
		name := names[pk.Response]
		io.String(&name)
		if pk.Response == packet.PackResponseSendPacks {
			protocol.FuncSlice(io, &pk.PacksToDownload, io.String)
		}
		return
	}
	response := uint8(pk.Response + 1)
	io.Uint8(&response)
	pk.Response = uint32(response) - 1
	count := uint16(len(pk.PacksToDownload))
	io.Uint16(&count)
	protocol.FuncSliceOfLen(io, uint32(count), &pk.PacksToDownload, io.String)
}
