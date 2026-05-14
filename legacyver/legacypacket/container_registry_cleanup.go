package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ContainerRegistryCleanup(io protocol.IO, pk *packet.ContainerRegistryCleanup) {
	protocol.FuncIOSlice(io, &pk.RemovedContainers, proto.MarshalFullContainerName)
}
