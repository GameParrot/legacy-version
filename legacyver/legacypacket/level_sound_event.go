package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func LevelSoundEvent(io protocol.IO, pk *packet.LevelSoundEvent) {
	if proto.IsProtoGTE(io, proto.ID1001) {
		io.String(&pk.SoundType)
	} else {
		legacySoundType(io, &pk.SoundType)
	}
	io.Vec3(&pk.Position)
	io.Varint32(&pk.ExtraData)
	io.String(&pk.EntityType)
	io.Bool(&pk.BabyMob)
	io.Bool(&pk.DisableRelativeVolume)
	if proto.IsProtoGTE(io, proto.ID786) {
		io.Int64(&pk.EntityUniqueID)
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.OptionalFunc(io, &pk.FireAtPosition, io.Vec3)
		}
	}
}
