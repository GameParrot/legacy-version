package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	CameraAunAssistPresetOperationSet = iota
	CameraAunAssistPresetOperationAddToExisting
)

// CameraAimAssistPresets is sent by the server to the client to provide a list of categories and presets
// that can be used when sending a CameraAimAssist packet or a CameraInstruction including aim assist.
type CameraAimAssistPresets struct {
	// Categories is a list of groups of categories which can be referenced by one of the Presets.
	Categories []protocol.CameraAimAssistCategory
	// Presets is a list of presets which define a base for how aim assist should behave
	Presets []protocol.CameraAimAssistPreset
	// Operation is the operation to perform with the presets. It is one of the constants above.
	Operation byte
}

// ID ...
func (*CameraAimAssistPresets) ID() uint32 {
	return packet.IDCameraAimAssistPresets
}

func (pk *CameraAimAssistPresets) Marshal(io protocol.IO) {
	protocol.Slice(io, &pk.Categories)
	protocol.Slice(io, &pk.Presets)
	if proto.IsProtoGTE(io, proto.ID776) {
		io.Uint8(&pk.Operation)
	}
}

// CameraAimAssistCategoryGroup is a group of categories which can be used by a CameraAimAssistPreset.
type CameraAimAssistCategoryGroup struct {
	// Identifier is the unique identifier of the group.
	Identifier string
	// Categories is a list of categories within this group.
	Categories []protocol.CameraAimAssistCategory
}

// Marshal encodes/decodes a CameraAimAssistCategoryGroup.
func (x *CameraAimAssistCategoryGroup) Marshal(r protocol.IO) {
	r.String(&x.Identifier)
	protocol.Slice(r, &x.Categories)
}
