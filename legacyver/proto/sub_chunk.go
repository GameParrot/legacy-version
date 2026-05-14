package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalSubChunkEntry(r protocol.IO, x *protocol.SubChunkEntry) {
	protocol.Single(r, &x.Offset)
	r.Uint8(&x.Result)
	if x.Result != protocol.SubChunkResultSuccessAllAir {
		r.ByteSlice(&x.RawPayload)
	}
	r.Uint8(&x.HeightMapType)
	if x.HeightMapType == protocol.HeightMapDataHasData {
		protocol.FuncSliceOfLen(r, 256, &x.HeightMapData, r.Int8)
	}
	if IsProtoGTE(r, ID818) {
		r.Uint8(&x.RenderHeightMapType)
		if x.RenderHeightMapType == protocol.HeightMapDataHasData {
			protocol.FuncSliceOfLen(r, 256, &x.RenderHeightMapData, r.Int8)
		}
	}
	r.Uint64(&x.BlobHash)
}

// SubChunkEntryNoCache encodes/decodes a SubChunkEntry assuming the blob cache is not enabled.
func SubChunkEntryNoCache(r protocol.IO, x *protocol.SubChunkEntry) {
	protocol.Single(r, &x.Offset)
	r.Uint8(&x.Result)
	r.ByteSlice(&x.RawPayload)
	r.Uint8(&x.HeightMapType)
	if x.HeightMapType == protocol.HeightMapDataHasData {
		protocol.FuncSliceOfLen(r, 256, &x.HeightMapData, r.Int8)
	}
	if IsProtoGTE(r, ID818) {
		r.Uint8(&x.RenderHeightMapType)
		if x.RenderHeightMapType == protocol.HeightMapDataHasData {
			protocol.FuncSliceOfLen(r, 256, &x.RenderHeightMapData, r.Int8)
		}
	}
}
