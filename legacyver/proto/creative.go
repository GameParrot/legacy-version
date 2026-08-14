package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalCreativeGroup(r protocol.IO, x *protocol.CreativeGroup) {
	if IsProtoGTE(r, ID2168) {
		r.Uint8(&x.Category)
	} else {
		protocol.IntegerFunc(&x.Category, r.Int32)
	}
	r.String(&x.Name)
	r.Item(&x.Icon)
}

func MarshalCreativeItem(r protocol.IO, x *protocol.CreativeItem) {
	r.Varuint32(&x.CreativeItemNetworkID)
	r.Item(&x.Item)
	if IsProtoGTE(r, ID776) {
		r.Varuint32(&x.GroupIndex)
	}
}
