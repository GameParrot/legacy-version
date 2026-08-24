package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	moveActorHasX = 1 << iota
	moveActorHasY
	moveActorHasZ
	moveActorHasRotX
	moveActorHasRotY
	moveActorHasRotYHead
	moveActorOnGround
	moveActorTeleport
	moveActorForceMove
)

func MoveActorDelta(io protocol.IO, pk *packet.MoveActorDelta) {
	io.Varuint64(&pk.EntityRuntimeID)
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.OptionalFunc(io, &pk.PositionX, io.Float32)
		protocol.OptionalFunc(io, &pk.PositionY, io.Float32)
		protocol.OptionalFunc(io, &pk.PositionZ, io.Float32)
		protocol.OptionalFunc(io, &pk.RotationX, io.ByteFloat)
		protocol.OptionalFunc(io, &pk.RotationY, io.ByteFloat)
		protocol.OptionalFunc(io, &pk.RotationYHead, io.ByteFloat)
		io.Bool(&pk.OnGround)
		io.Bool(&pk.ForceMove)
		io.Bool(&pk.ForceMoveLocalEntity)
		io.Bool(&pk.ForceCompletion)
		if proto.IsProtoGTE(io, proto.ID2192) {
			io.Varuint64(&pk.Ticks)
		}
		return
	}
	flags := uint16(0)
	setFlag := func(optional protocol.Optional[float32], flag uint16) {
		if _, ok := optional.Value(); ok {
			flags |= flag
		}
	}
	setFlag(pk.PositionX, moveActorHasX)
	setFlag(pk.PositionY, moveActorHasY)
	setFlag(pk.PositionZ, moveActorHasZ)
	setFlag(pk.RotationX, moveActorHasRotX)
	setFlag(pk.RotationY, moveActorHasRotY)
	setFlag(pk.RotationYHead, moveActorHasRotYHead)
	if pk.OnGround {
		flags |= moveActorOnGround
	}
	if pk.ForceCompletion {
		flags |= moveActorTeleport
	}
	if pk.ForceMove {
		flags |= moveActorForceMove
	}
	io.Uint16(&flags)
	legacyOptionalFloat(io, &pk.PositionX, flags&moveActorHasX != 0, io.Float32)
	legacyOptionalFloat(io, &pk.PositionY, flags&moveActorHasY != 0, io.Float32)
	legacyOptionalFloat(io, &pk.PositionZ, flags&moveActorHasZ != 0, io.Float32)
	legacyOptionalFloat(io, &pk.RotationX, flags&moveActorHasRotX != 0, io.ByteFloat)
	legacyOptionalFloat(io, &pk.RotationY, flags&moveActorHasRotY != 0, io.ByteFloat)
	legacyOptionalFloat(io, &pk.RotationYHead, flags&moveActorHasRotYHead != 0, io.ByteFloat)
	pk.OnGround = flags&moveActorOnGround != 0
	pk.ForceCompletion = flags&moveActorTeleport != 0
	pk.ForceMove = flags&moveActorForceMove != 0
}

func legacyOptionalFloat(io protocol.IO, x *protocol.Optional[float32], present bool, f func(*float32)) {
	if !present {
		return
	}
	v, _ := x.Value()
	f(&v)
	*x = protocol.Option(v)
}
