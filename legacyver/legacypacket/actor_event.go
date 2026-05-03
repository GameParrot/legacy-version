package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// ActorEvent is sent by the server when a particular event happens that has to do with an entity. Some of
// these events are entity-specific, for example a wolf shaking itself dry, but others are used for each
// entity, such as dying.
type ActorEvent struct {
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// EventType is the ID of the event to be called. It is one of the constants that can be found above.
	EventType byte
	// EventData is optional data associated with a particular event. The data has a different function for
	// different events, however most events don't use this field at all.
	EventData int32
}

// ID ...
func (*ActorEvent) ID() uint32 {
	return packet.IDActorEvent
}

func (pk *ActorEvent) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint8(&pk.EventType)
	io.Varint32(&pk.EventData)
	if proto.IsProtoGTE(io, proto.ID974) {
		var zero bool
		io.Bool(&zero)
	}
}
