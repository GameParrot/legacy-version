package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalDataStoreUpdate(r protocol.IO, x *protocol.DataStoreUpdate) {
	r.String(&x.DataStoreName)
	r.String(&x.Property)
	r.String(&x.Path)
	if IsProtoLT(r, ID1001) {
		r.Uint32(&x.ControlType)
	} else {
		r.Varuint32(&x.ControlType)
	}
	switch x.ControlType {
	case protocol.DataStoreControlDouble:
		r.Float64(&x.DoubleValue)
	case protocol.DataStoreControlBoolean:
		r.Bool(&x.BoolValue)
	case protocol.DataStoreControlString:
		r.String(&x.StringValue)
	default:
		r.UnknownEnumOption(x.ControlType, "data store control type")
	}
	r.Uint32(&x.PropertyUpdateCount)
	r.Uint32(&x.PathUpdateCount)
}

func MarshalDataStoreChangeEntry(r protocol.IO, x *protocol.DataStoreChangeEntry) {
	if IsProtoLT(r, ID1001) {
		r.Uint32(&x.ChangeType)
	} else {
		r.Varuint32(&x.ChangeType)
	}
	switch x.ChangeType {
	case protocol.DataStoreChangeTypeUpdate:
		MarshalDataStoreUpdate(r, &x.Update)
	case protocol.DataStoreChangeTypeChange:
		r.String(&x.Change.DataStoreName)
		r.String(&x.Change.Property)
		r.Uint32(&x.Change.UpdateCount)
		MarshalDataStorePropertyValue(r, &x.Change.NewValue)
	case protocol.DataStoreChangeTypeRemoval:
		r.String(&x.Removal.DataStoreName)
	default:
		r.UnknownEnumOption(x.ChangeType, "data store change type")
	}
}

func MarshalDataStorePropertyValue(r protocol.IO, x *protocol.DataStorePropertyValue) {
	r.Int32(&x.Type)
	switch x.Type {
	case protocol.DataStorePropertyTypeNone:
	case protocol.DataStorePropertyTypeBool:
		r.Bool(&x.BoolValue)
	case protocol.DataStorePropertyTypeInt64:
		r.Int64(&x.Int64Value)
	case protocol.DataStorePropertyTypeDouble:
		if IsProtoLT(r, ID1001) {
			r.UnknownEnumOption(x.Type, "data store property type")
		} else {
			r.Float64(&x.DoubleValue)
		}
	case protocol.DataStorePropertyTypeString:
		r.String(&x.StringValue)
	case protocol.DataStorePropertyTypeList:
		if IsProtoLT(r, ID1001) {
			r.UnknownEnumOption(x.Type, "data store property type")
		} else {
			protocol.FuncIOSlice(r, &x.ListValue, MarshalDataStorePropertyValue)
		}
	case protocol.DataStorePropertyTypeMap:
		protocol.FuncSlice(r, &x.MapValue, func(entry *protocol.DataStoreMapEntry) {
			r.String(&entry.Key)
			MarshalDataStorePropertyValue(r, &entry.Value)
		})
	default:
		r.UnknownEnumOption(x.Type, "data store property type")
	}
}
