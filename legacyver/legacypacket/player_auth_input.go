package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerAuthInput(io protocol.IO, pk *packet.PlayerAuthInput) {
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Vec3(&pk.Position)
	io.Vec2(&pk.MoveVector)
	io.Float32(&pk.HeadYaw)
	if proto.IsProtoGTE(io, proto.ID766) {
		io.Bitset(&pk.InputData, packet.PlayerAuthInputBitsetSize)
	} else {
		io.Bitset(&pk.InputData, 64)
	}
	io.Varuint32(&pk.InputMode)
	io.Varuint32(&pk.PlayMode)
	io.Varuint32(&pk.InteractionModel)
	if proto.IsProtoGTE(io, proto.ID748) {
		io.Float32(&pk.InteractPitch)
		io.Float32(&pk.InteractYaw)
	}
	io.Varuint64(&pk.Tick)
	io.Vec3(&pk.Delta)

	if pk.InputData.Load(packet.InputFlagPerformItemInteraction) {
		proto.PlayerInventoryAction(io, &pk.ItemInteractionData)
	}

	if pk.InputData.Load(packet.InputFlagPerformItemStackRequest) {
		proto.MarshalItemStackRequest(io, &pk.ItemStackRequest)
	}

	if pk.InputData.Load(packet.InputFlagPerformBlockActions) {
		protocol.SliceVarint32Length(io, &pk.BlockActions)
	}

	if pk.InputData.Load(packet.InputFlagClientPredictedVehicle) {
		if proto.IsProtoGTE(io, proto.ID662) {
			io.Vec2(&pk.VehicleRotation)
		}
		io.Varint64(&pk.ClientPredictedVehicle)
	}

	io.Vec2(&pk.AnalogueMoveVector)
	if proto.IsProtoGTE(io, proto.ID748) {
		io.Vec3(&pk.CameraOrientation)
	}

	if proto.IsProtoGTE(io, proto.ID766) {
		io.Vec2(&pk.RawMoveVector)
	}
}
