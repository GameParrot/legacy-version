package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalFullContainerName(r protocol.IO, x *protocol.FullContainerName) {
	r.Uint8(&x.ContainerID)
	if IsProtoGTE(r, ID729) {
		protocol.OptionalFunc(r, &x.DynamicContainerID, r.Uint32)
	} else if IsProtoGTE(r, ID712) {
		dynamicContainerID, _ := x.DynamicContainerID.Value()
		r.Uint32(&dynamicContainerID)
		if dynamicContainerID != 0 {
			x.DynamicContainerID = protocol.Option(dynamicContainerID)
		} else {
			x.DynamicContainerID = protocol.Optional[uint32]{}
		}
	}
}
