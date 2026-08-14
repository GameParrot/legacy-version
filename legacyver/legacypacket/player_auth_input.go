package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const legacyPlayerAuthInputBitsetSize = 65

func PlayerAuthInput(io protocol.IO, pk *packet.PlayerAuthInput) {
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Vec3(&pk.Position)
	io.Vec2(&pk.MoveVector)
	io.Float32(&pk.HeadYaw)
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.InputFlagList(io, &pk.InputData, packet.InputFlagCount)
	} else {
		size := 64
		if proto.IsProtoGTE(io, proto.ID766) {
			size = legacyPlayerAuthInputBitsetSize
		}
		legacyInputFlags(io, &pk.InputData, size)
	}
	io.Varuint32(&pk.InputMode)
	io.Varuint32(&pk.PlayMode)
	if proto.IsProtoGTE(io, proto.ID2168) {
		io.Varint32(&pk.InteractionModel)
	} else {
		interactionModel := uint32(pk.InteractionModel)
		io.Varuint32(&interactionModel)
		pk.InteractionModel = int32(interactionModel)
	}
	if proto.IsProtoGTE(io, proto.ID748) {
		io.Float32(&pk.InteractPitch)
		io.Float32(&pk.InteractYaw)
	}
	io.Varuint64(&pk.Tick)
	io.Vec3(&pk.Delta)
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.DoubleOptionalFunc(io, &pk.ItemInteractionData, func(x *protocol.UseItemTransactionData) {
			proto.PlayerInventoryAction(io, x)
		})
	} else if pk.InputData.Load(packet.InputFlagPerformItemInteraction) {
		x, _ := pk.ItemInteractionData.Value()
		proto.PlayerInventoryAction(io, &x)
		pk.ItemInteractionData = protocol.Option(x)
	}
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.DoubleOptionalFunc(io, &pk.ItemStackRequest, func(x *protocol.ItemStackRequest) {
			proto.MarshalItemStackRequest(io, x)
		})
	} else if pk.InputData.Load(packet.InputFlagPerformItemStackRequest) {
		x, _ := pk.ItemStackRequest.Value()
		proto.MarshalItemStackRequest(io, &x)
		pk.ItemStackRequest = protocol.Option(x)
	}
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.DoubleOptionalFunc(io, &pk.BlockActions, func(x *[]protocol.PlayerBlockAction) {
			protocol.FuncIOSlice(io, x, proto.MarshalPlayerBlockAction)
		})
	} else if pk.InputData.Load(packet.InputFlagPerformBlockActions) {
		x, _ := pk.BlockActions.Value()
		count := int32(len(x))
		io.Varint32(&count)
		protocol.FuncIOSliceOfLen(io, uint32(count), &x, proto.MarshalPlayerBlockAction)
		pk.BlockActions = protocol.Option(x)
	}
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.DoubleOptionalFunc(io, &pk.VehicleRotation, io.Vec2)
		protocol.DoubleOptionalFunc(io, &pk.ClientPredictedVehicle, io.Varint64)
	} else if pk.InputData.Load(packet.InputFlagClientPredictedVehicle) {
		if proto.IsProtoGTE(io, proto.ID662) {
			x, _ := pk.VehicleRotation.Value()
			io.Vec2(&x)
			pk.VehicleRotation = protocol.Option(x)
		}
		x, _ := pk.ClientPredictedVehicle.Value()
		io.Varint64(&x)
		pk.ClientPredictedVehicle = protocol.Option(x)
	}
	io.Vec2(&pk.AnalogueMoveVector)
	if proto.IsProtoGTE(io, proto.ID748) {
		io.Vec3(&pk.CameraOrientation)
	}
	if proto.IsProtoGTE(io, proto.ID766) {
		io.Vec2(&pk.RawMoveVector)
	}
}

func legacyInputFlags(io protocol.IO, flags *protocol.InputFlags, size int) {
	bits := protocol.NewBitset(size)
	if !proto.IsReader(io) {
		for i := 0; i < size && i < flags.Len(); i++ {
			if flags.Load(i) {
				bits.Set(i)
			}
		}
	}
	io.Bitset(&bits, size)
	if proto.IsReader(io) {
		*flags = protocol.NewInputFlags(size)
		for i := 0; i < size; i++ {
			if bits.Load(i) {
				flags.Set(i)
			}
		}
	}
}
