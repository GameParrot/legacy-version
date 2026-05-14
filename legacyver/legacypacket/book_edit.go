package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func BookEdit(io protocol.IO, pk *packet.BookEdit) {
	if proto.IsProtoGTE(io, proto.ID924) {
		io.Varint32(&pk.InventorySlot)
		io.Varuint32(&pk.ActionType)
		switch pk.ActionType {
		case packet.BookActionReplacePage, packet.BookActionAddPage:
			io.Varint32(&pk.PageNumber)
			io.String(&pk.Text)
			io.String(&pk.PhotoName)
		case packet.BookActionDeletePage:
			io.Varint32(&pk.PageNumber)
		case packet.BookActionSwapPages:
			io.Varint32(&pk.PageNumber)
			io.Varint32(&pk.SecondaryPageNumber)
		case packet.BookActionSign:
			io.String(&pk.Title)
			io.String(&pk.Author)
			io.String(&pk.XUID)
		default:
			io.UnknownEnumOption(pk.ActionType, "book edit action type")
		}
		return
	}

	actionType := byte(pk.ActionType)
	io.Uint8(&actionType)
	pk.ActionType = uint32(actionType)

	inventorySlot := byte(pk.InventorySlot)
	io.Uint8(&inventorySlot)
	pk.InventorySlot = int32(inventorySlot)

	switch pk.ActionType {
	case packet.BookActionReplacePage, packet.BookActionAddPage:
		pageNumber := byte(pk.PageNumber)
		io.Uint8(&pageNumber)
		pk.PageNumber = int32(pageNumber)
		io.String(&pk.Text)
		io.String(&pk.PhotoName)
	case packet.BookActionDeletePage:
		pageNumber := byte(pk.PageNumber)
		io.Uint8(&pageNumber)
		pk.PageNumber = int32(pageNumber)
	case packet.BookActionSwapPages:
		pageNumber := byte(pk.PageNumber)
		io.Uint8(&pageNumber)
		pk.PageNumber = int32(pageNumber)
		secondaryPageNumber := byte(pk.SecondaryPageNumber)
		io.Uint8(&secondaryPageNumber)
		pk.SecondaryPageNumber = int32(secondaryPageNumber)
	case packet.BookActionSign:
		io.String(&pk.Title)
		io.String(&pk.Author)
		io.String(&pk.XUID)
	default:
		io.UnknownEnumOption(pk.ActionType, "book edit action type")
	}
}
