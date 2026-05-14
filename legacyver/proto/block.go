package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalBlockChangeEntry(r protocol.IO, x *protocol.BlockChangeEntry) {
	IOUBlockPos(r, &x.BlockPos)
	r.Varuint32(&x.BlockRuntimeID)
	r.Varuint32(&x.Flags)
	r.Varuint64(&x.SyncedUpdateEntityUniqueID)
	r.Varuint32(&x.SyncedUpdateType)
}
