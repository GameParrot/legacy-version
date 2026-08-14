package proto

import (
	"fmt"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalGameRule(r protocol.IO, x *protocol.GameRule, legacyInteger bool) {
	if IsProtoGTE(r, ID2168) {
		r.GameRule(x)
		return
	}
	r.String(&x.Name)
	r.Bool(&x.CanBeModifiedByPlayer)
	var kind uint32
	if !IsReader(r) {
		switch x.Value.(type) {
		case bool:
			kind = 1
		case uint32:
			kind = 2
		case float32:
			kind = 3
		default:
			r.UnknownEnumOption(fmt.Sprintf("%T", x.Value), "game rule type")
			return
		}
	}
	r.Varuint32(&kind)
	switch kind {
	case 1:
		v, _ := x.Value.(bool)
		r.Bool(&v)
		x.Value = v
	case 2:
		v, _ := x.Value.(uint32)
		if legacyInteger {
			r.Varuint32(&v)
		} else {
			r.Uint32(&v)
		}
		x.Value = v
	case 3:
		v, _ := x.Value.(float32)
		r.Float32(&v)
		x.Value = v
	default:
		r.UnknownEnumOption(kind, "game rule type")
	}
}
