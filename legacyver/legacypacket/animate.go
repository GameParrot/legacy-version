package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// Animate is sent by the server to send a player animation from one player to all viewers of that player. It
// is used for a couple of actions, such as arm swimming and critical hits.
type Animate struct {
	// ActionType is the ID of the animation action to execute. It is one of the action type constants that
	// may be found above.
	ActionType int32
	// EntityRuntimeID is the runtime ID of the player that the animation should be played upon. The runtime
	// ID is unique for each world session, and entities are generally identified in packets using this
	// runtime ID.
	EntityRuntimeID uint64
	// Data ...
	Data float32
	// RowingTime is the time for rowing actions.
	RowingTime float32
}

// ID ...
func (*Animate) ID() uint32 {
	return packet.IDAnimate
}

func (pk *Animate) Marshal(io protocol.IO) {
	io.Varint32(&pk.ActionType)
	io.Varuint64(&pk.EntityRuntimeID)
	if proto.IsProtoGTE(io, proto.ID859) {
		io.Float32(&pk.Data)
		if pk.ActionType == packet.AnimateActionRowLeft || pk.ActionType == packet.AnimateActionRowRight {
			io.Float32(&pk.RowingTime)
		}
	} else {
		if pk.ActionType&0x80 != 0 {
			io.Float32(&pk.RowingTime)
		}
	}
}
