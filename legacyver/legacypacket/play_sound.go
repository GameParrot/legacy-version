package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlaySound(io protocol.IO, pk *packet.PlaySound) {
	io.String(&pk.SoundName)
	io.SoundPos(&pk.Position)
	io.Float32(&pk.Volume)
	io.Float32(&pk.Pitch)
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Varint32(&pk.LoopCount)
	}
	if proto.IsProtoGTE(io, proto.ID975) {
		protocol.OptionalFunc(io, &pk.Handle, io.Uint64)
	}
}
