package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalSubChunkEntry(r protocol.IO, x *protocol.SubChunkEntry) {
	MarshalSubChunkOffset(r, &x.Offset)
	r.Uint8(&x.Result)
	if IsProtoGTE(r, ID2168) {
		protocol.OptionalFunc(r, &x.RawPayload, r.ByteSlice)
	} else if x.Result != protocol.SubChunkResultSuccessAllAir {
		payload, _ := x.RawPayload.Value()
		r.ByteSlice(&payload)
		x.RawPayload = protocol.Option(payload)
	}
	r.Uint8(&x.HeightMapType)
	if IsProtoGTE(r, ID2168) {
		protocol.OptionalFunc(r, &x.HeightMapData, func(data *[]int8) {
			protocol.FuncSliceOfLen(r, 256, data, r.Int8)
		})
	} else if x.HeightMapType == protocol.HeightMapDataHasData {
		data, _ := x.HeightMapData.Value()
		protocol.FuncSliceOfLen(r, 256, &data, r.Int8)
		x.HeightMapData = protocol.Option(data)
	}
	if IsProtoGTE(r, ID818) {
		r.Uint8(&x.RenderHeightMapType)
		if IsProtoGTE(r, ID2168) {
			protocol.OptionalFunc(r, &x.RenderHeightMapData, func(data *[]int8) {
				protocol.FuncSliceOfLen(r, 256, data, r.Int8)
			})
		} else if x.RenderHeightMapType == protocol.HeightMapDataHasData {
			data, _ := x.RenderHeightMapData.Value()
			protocol.FuncSliceOfLen(r, 256, &data, r.Int8)
			x.RenderHeightMapData = protocol.Option(data)
		}
	}
	if IsProtoGTE(r, ID2168) {
		protocol.OptionalFunc(r, &x.BlobHash, r.Uint64)
	} else {
		blobHash, _ := x.BlobHash.Value()
		r.Uint64(&blobHash)
		x.BlobHash = protocol.Option(blobHash)
	}
}

// SubChunkEntryNoCache encodes/decodes a SubChunkEntry assuming the blob cache is not enabled.
func SubChunkEntryNoCache(r protocol.IO, x *protocol.SubChunkEntry) {
	MarshalSubChunkOffset(r, &x.Offset)
	r.Uint8(&x.Result)
	payload, _ := x.RawPayload.Value()
	r.ByteSlice(&payload)
	x.RawPayload = protocol.Option(payload)
	r.Uint8(&x.HeightMapType)
	if x.HeightMapType == protocol.HeightMapDataHasData {
		data, _ := x.HeightMapData.Value()
		protocol.FuncSliceOfLen(r, 256, &data, r.Int8)
		x.HeightMapData = protocol.Option(data)
	}
	if IsProtoGTE(r, ID818) {
		r.Uint8(&x.RenderHeightMapType)
		if x.RenderHeightMapType == protocol.HeightMapDataHasData {
			data, _ := x.RenderHeightMapData.Value()
			protocol.FuncSliceOfLen(r, 256, &data, r.Int8)
			x.RenderHeightMapData = protocol.Option(data)
		}
	}
}

func MarshalSubChunkOffset(r protocol.IO, x *protocol.SubChunkOffset) {
	r.Int8(&x[0])
	r.Int8(&x[1])
	r.Int8(&x[2])
}
