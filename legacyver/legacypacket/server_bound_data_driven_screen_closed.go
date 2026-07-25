package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ServerBoundDataDrivenScreenClosed(io protocol.IO, pk *packet.ServerBoundDataDrivenScreenClosed) {
	io.Uint32(&pk.FormID)
	if proto.IsProtoGTE(io, proto.ID1001) {
		io.String(&pk.CloseReason)
		return
	}
	if proto.IsReader(io) {
		var reason uint8
		io.Uint8(&reason)
		pk.CloseReason = dataDrivenScreenCloseReasonToString(io, reason)
		return
	}
	reason := dataDrivenScreenCloseReasonFromString(io, pk.CloseReason)
	io.Uint8(&reason)
}

func dataDrivenScreenCloseReasonToString(io protocol.IO, reason uint8) string {
	switch reason {
	case 0:
		return packet.DataDrivenScreenCloseReasonProgrammaticClose
	case 1:
		return packet.DataDrivenScreenCloseReasonProgrammaticCloseAll
	case 2:
		return packet.DataDrivenScreenCloseReasonClientCanceled
	case 3:
		return packet.DataDrivenScreenCloseReasonUserBusy
	case 4:
		return packet.DataDrivenScreenCloseReasonInvalidForm
	default:
		io.UnknownEnumOption(reason, "data driven screen close reason")
		return ""
	}
}

func dataDrivenScreenCloseReasonFromString(io protocol.IO, reason string) uint8 {
	switch reason {
	case packet.DataDrivenScreenCloseReasonProgrammaticClose:
		return 0
	case packet.DataDrivenScreenCloseReasonProgrammaticCloseAll:
		return 1
	case packet.DataDrivenScreenCloseReasonClientCanceled:
		return 2
	case packet.DataDrivenScreenCloseReasonUserBusy:
		return 3
	case packet.DataDrivenScreenCloseReasonInvalidForm:
		return 4
	default:
		io.UnknownEnumOption(reason, "data driven screen close reason")
		return 0
	}
}
