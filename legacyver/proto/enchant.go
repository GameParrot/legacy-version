package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalEnchantmentOption(r protocol.IO, x *protocol.EnchantmentOption) {
	if IsProtoGTE(r, ID975) {
		r.Uint8(&x.Cost)
	} else {
		costUint32 := uint32(x.Cost)
		r.Varuint32(&costUint32)
		x.Cost = uint8(costUint32)
	}
	protocol.Single(r, &x.Enchantments)
	r.String(&x.Name)
	r.Varuint32(&x.RecipeNetworkID)
}
