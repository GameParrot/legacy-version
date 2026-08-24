package proto

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

func MarshalPackSetting(r protocol.IO, x *protocol.PackSetting) {
	r.String(&x.Name)
	if IsReader(r) {
		var valueType uint32
		r.Varuint32(&valueType)
		switch valueType {
		case protocol.PackSettingTypeFloat:
			var value float32
			r.Float32(&value)
			x.Value = value
		case protocol.PackSettingTypeBool:
			var value bool
			r.Bool(&value)
			x.Value = value
		case protocol.PackSettingTypeString:
			var value string
			r.String(&value)
			x.Value = value
		case protocol.PackSettingTypeStringList:
			if IsProtoLT(r, ID2192) {
				r.UnknownEnumOption(valueType, "pack setting")
				return
			}
			var value []string
			protocol.FuncSlice(r, &value, r.String)
			x.Value = value
		default:
			r.UnknownEnumOption(valueType, "pack setting")
		}
		return
	}
	var valueType uint32
	switch value := x.Value.(type) {
	case float32:
		valueType = protocol.PackSettingTypeFloat
		r.Varuint32(&valueType)
		r.Float32(&value)
	case bool:
		valueType = protocol.PackSettingTypeBool
		r.Varuint32(&valueType)
		r.Bool(&value)
	case string:
		valueType = protocol.PackSettingTypeString
		r.Varuint32(&valueType)
		r.String(&value)
	case []string:
		if IsProtoLT(r, ID2192) {
			r.UnknownEnumOption(fmt.Sprintf("%T", x.Value), "pack setting")
			return
		}
		valueType = protocol.PackSettingTypeStringList
		r.Varuint32(&valueType)
		protocol.FuncSlice(r, &value, r.String)
	default:
		r.UnknownEnumOption(fmt.Sprintf("%T", x.Value), "pack setting")
	}
}

func MarshalTexturePackInfo(r protocol.IO, x *protocol.TexturePackInfo) {
	if IsProtoGTE(r, ID766) {
		r.UUID(&x.UUID)
	} else {
		if IsReader(r) {
			uuidStr := ""
			r.String(&uuidStr)
			x.UUID = uuid.MustParse(uuidStr)
		} else {
			uuidStr := x.UUID.String()
			r.String(&uuidStr)
		}
	}
	r.String(&x.Version)
	r.Uint64(&x.Size)
	r.String(&x.ContentKey)
	r.String(&x.SubPackName)
	r.String(&x.ContentIdentity)
	r.Bool(&x.HasScripts)
	if IsProtoGTE(r, ID712) {
		r.Bool(&x.AddonPack)
	}
	r.Bool(&x.RTXEnabled)
	if IsProtoGTE(r, ID748) {
		r.String(&x.DownloadURL)
	}
}

// BehaviourPackInfo represents a behaviour pack's info sent over network. It holds information about the
// behaviour pack such as its name, description and version.
type BehaviourPackInfo struct {
	// UUID is the UUID of the behaviour pack. Each behaviour pack downloaded must have a different UUID in
	// order for the client to be able to handle them properly.
	UUID string
	// Version is the version of the behaviour pack. The client will cache behaviour packs sent by the server as
	// long as they carry the same version. Sending a behaviour pack with a different version than previously
	// will force the client to re-download it.
	Version string
	// Size is the total size in bytes that the behaviour pack occupies. This is the size of the compressed
	// archive (zip) of the behaviour pack.
	Size uint64
	// ContentKey is the key used to decrypt the behaviour pack if it is encrypted. This is generally the case
	// for marketplace behaviour packs.
	ContentKey string
	// SubPackName ...
	SubPackName string
	// ContentIdentity ...
	ContentIdentity string
	// HasScripts specifies if the behaviour packs has any scripts in it. A client will only download the
	// behaviour pack if it supports scripts, which, up to 1.11, only includes Windows 10.
	HasScripts bool
	// AddonPack specifies if the texture pack is from an addon.
	AddonPack bool
}

// Marshal encodes/decodes a BehaviourPackInfo.
func (x *BehaviourPackInfo) Marshal(r protocol.IO) {
	r.String(&x.UUID)
	r.String(&x.Version)
	r.Uint64(&x.Size)
	r.String(&x.ContentKey)
	r.String(&x.SubPackName)
	r.String(&x.ContentIdentity)
	r.Bool(&x.HasScripts)
	if IsProtoGTE(r, ID712) {
		r.Bool(&x.AddonPack)
	}
}
