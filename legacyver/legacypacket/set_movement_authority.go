package legacypacket

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const IDSetMovementAuthority uint32 = 319

// SetMovementAuthority is the packet removed after the legacy protocols. It is
// retained so old connections can still decode and round-trip packet ID 319.
type SetMovementAuthority struct {
	MovementType byte
}

func (*SetMovementAuthority) ID() uint32 { return IDSetMovementAuthority }

func (pk *SetMovementAuthority) Marshal(io protocol.IO) {
	io.Uint8(&pk.MovementType)
}
