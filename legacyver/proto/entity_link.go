package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// Marshal encodes/decodes a single entity link.
func MarshalEntityLink(r protocol.IO, x *protocol.EntityLink) {
	r.Varint64(&x.RiddenEntityUniqueID)
	r.Varint64(&x.RiderEntityUniqueID)
	r.Uint8(&x.Type)
	r.Bool(&x.Immediate)
	r.Bool(&x.RiderInitiated)
	if IsProtoGTE(r, ID712) {
		r.Float32(&x.VehicleAngularVelocity)
	}
}
