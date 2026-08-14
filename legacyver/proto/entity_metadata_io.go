package proto

import (
	"reflect"
	"sort"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// MarshalEntityMetadata encodes each metadata entry using the layout of the
// selected protocol. Protocol 2168 added a second byte-sized copy of the type
// discriminator after the existing varuint discriminator.
func MarshalEntityMetadata(r protocol.IO, metadata *protocol.EntityMetadata) {
	if IsReader(r) {
		*metadata = protocol.EntityMetadata{}
	}
	count := uint32(len(*metadata))
	r.Varuint32(&count)
	if IsReader(r) {
		for range count {
			var key, dataType uint32
			r.Varuint32(&key)
			r.Varuint32(&dataType)
			marshalEntityMetadataLegacyType(r, dataType)
			(*metadata)[key] = readEntityMetadataValue(r, dataType)
		}
		return
	}
	keys := make([]int, 0, count)
	for key := range *metadata {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)
	for _, rawKey := range keys {
		key := uint32(rawKey)
		r.Varuint32(&key)
		writeEntityMetadataValue(r, (*metadata)[key])
	}
}

func marshalEntityMetadataLegacyType(r protocol.IO, dataType uint32) {
	if IsProtoGTE(r, ID2168) {
		legacyType := byte(dataType)
		r.Uint8(&legacyType)
		if uint32(legacyType) != dataType {
			r.InvalidValue(legacyType, "entity metadata type", "does not match primary discriminator")
		}
	}
}

func writeEntityMetadataValue(r protocol.IO, value any) {
	writeType := func(dataType uint32) {
		r.Varuint32(&dataType)
		marshalEntityMetadataLegacyType(r, dataType)
	}
	switch value := value.(type) {
	case byte:
		writeType(protocol.EntityDataTypeByte)
		r.Uint8(&value)
	case int16:
		writeType(protocol.EntityDataTypeInt16)
		r.Int16(&value)
	case int32:
		writeType(protocol.EntityDataTypeInt32)
		r.Varint32(&value)
	case float32:
		writeType(protocol.EntityDataTypeFloat32)
		r.Float32(&value)
	case string:
		writeType(protocol.EntityDataTypeString)
		r.String(&value)
	case map[string]any:
		writeType(protocol.EntityDataTypeCompoundTag)
		r.NBT(&value, nbt.NetworkLittleEndian)
	case protocol.BlockPos:
		writeType(protocol.EntityDataTypeBlockPos)
		r.BlockPos(&value)
	case int64:
		writeType(protocol.EntityDataTypeInt64)
		r.Varint64(&value)
	case mgl32.Vec3:
		writeType(protocol.EntityDataTypeVec3)
		r.Vec3(&value)
	default:
		r.UnknownEnumOption(reflect.TypeOf(value), "entity metadata")
	}
}

func readEntityMetadataValue(r protocol.IO, dataType uint32) any {
	switch dataType {
	case protocol.EntityDataTypeByte:
		var value byte
		r.Uint8(&value)
		return value
	case protocol.EntityDataTypeInt16:
		var value int16
		r.Int16(&value)
		return value
	case protocol.EntityDataTypeInt32:
		var value int32
		r.Varint32(&value)
		return value
	case protocol.EntityDataTypeFloat32:
		var value float32
		r.Float32(&value)
		return value
	case protocol.EntityDataTypeString:
		var value string
		r.String(&value)
		return value
	case protocol.EntityDataTypeCompoundTag:
		var value map[string]any
		r.NBT(&value, nbt.NetworkLittleEndian)
		return value
	case protocol.EntityDataTypeBlockPos:
		var value protocol.BlockPos
		r.BlockPos(&value)
		return value
	case protocol.EntityDataTypeInt64:
		var value int64
		r.Varint64(&value)
		return value
	case protocol.EntityDataTypeVec3:
		var value mgl32.Vec3
		r.Vec3(&value)
		return value
	default:
		r.UnknownEnumOption(dataType, "entity metadata")
		return nil
	}
}
