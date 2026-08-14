package legacypacket

import (
	legacyproto "github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientBoundAttributeLayerSync(r protocol.IO, pk *packet.ClientBoundAttributeLayerSync) {
	r.Varuint32(&pk.PayloadType)
	switch pk.PayloadType {
	case protocol.AttributeLayerPayloadTypeUpdateLayers:
		protocol.FuncIOSlice(r, &pk.Layers, legacyproto.MarshalAttributeLayerData)
	case protocol.AttributeLayerPayloadTypeUpdateSettings:
		r.String(&pk.LayerName)
		r.Varint32(&pk.DimensionID)
		legacyproto.MarshalAttributeLayerSettings(r, &pk.Settings)
	case protocol.AttributeLayerPayloadTypeUpdateEnvironment:
		r.String(&pk.LayerName)
		r.Varint32(&pk.DimensionID)
		protocol.FuncIOSlice(r, &pk.EnvironmentAttributes, legacyproto.MarshalEnvironmentAttributeData)
	case protocol.AttributeLayerPayloadTypeRemoveEnvironment:
		r.String(&pk.LayerName)
		r.Varint32(&pk.DimensionID)
		protocol.FuncSlice(r, &pk.RemoveAttributeNames, r.String)
	default:
		r.UnknownEnumOption(pk.PayloadType, "attribute layer payload type")
	}
}
