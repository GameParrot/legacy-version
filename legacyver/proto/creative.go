package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalCreativeItem(r protocol.IO, x *protocol.CreativeItem) {
	r.Varuint32(&x.CreativeItemNetworkID)
	r.Item(&x.Item)
	if IsProtoGTE(r, ID776) {
		r.Varuint32(&x.GroupIndex)
	}
}
