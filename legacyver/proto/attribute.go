package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalAttribute(r protocol.IO, x *protocol.Attribute) {
	r.Float32(&x.Min)
	r.Float32(&x.Max)
	r.Float32(&x.Value)
	if IsProtoGTE(r, ID729) {
		r.Float32(&x.DefaultMin)
		r.Float32(&x.DefaultMax)
	}
	r.Float32(&x.Default)
	r.String(&x.Name)
	protocol.Slice(r, &x.Modifiers)
}
