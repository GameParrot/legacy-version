package proto

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// LegacyItemRegistryEntry is an item sent in the StartGame item table. It holds a name and a legacy ID, which is used to
// point back to that name.
type LegacyItemRegistryEntry struct {
	// Name if the name of the item, which is a name like 'minecraft:stick'.
	Name string
	// RuntimeID is the ID that is used to identify the item over network. After sending all items in the
	// StartGame packet, items will then be identified using these numerical IDs.
	RuntimeID int16
	// ComponentBased specifies if the item was created using components, meaning the item is a custom item.
	ComponentBased bool
}

// Marshal encodes/decodes an LegacyItemRegistryEntry.
func (x *LegacyItemRegistryEntry) Marshal(r protocol.IO) {
	r.String(&x.Name)
	r.Int16(&x.RuntimeID)
	r.Bool(&x.ComponentBased)
}

// FromLatest ...
func (x *LegacyItemRegistryEntry) FromLatest(y protocol.ItemEntry) LegacyItemRegistryEntry {
	x.Name = y.Name
	x.RuntimeID = y.RuntimeID
	x.ComponentBased = y.ComponentBased
	return *x
}

// ToLatest ...
func (x *LegacyItemRegistryEntry) ToLatest() protocol.ItemEntry {
	return protocol.ItemEntry{
		Name:           x.Name,
		RuntimeID:      x.RuntimeID,
		ComponentBased: x.ComponentBased,
		Version:        2,
	}
}

func MarshalItemEntry(r protocol.IO, x *protocol.ItemEntry) {
	r.String(&x.Name)
	if IsProtoGTE(r, ID776) {
		r.Int16(&x.RuntimeID)
		r.Bool(&x.ComponentBased)
		r.Varint32(&x.Version)
	}
	r.NBT(&x.Data, nbt.NetworkLittleEndian)
}
