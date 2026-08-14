package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func PlayerArmourDamage(io protocol.IO, pk *packet.PlayerArmourDamage) {
	if proto.IsProtoLT(io, proto.ID844) {
		var bitset uint8
		slots := []int32{packet.PlayerArmourDamageFlagHelmet, packet.PlayerArmourDamageFlagChestplate, packet.PlayerArmourDamageFlagLeggings, packet.PlayerArmourDamageFlagBoots}
		if proto.IsProtoGTE(io, proto.ID712) {
			slots = append(slots, packet.PlayerArmourDamageFlagBody)
		}
		if proto.IsReader(io) {
			io.Uint8(&bitset)
			pk.List = make([]protocol.PlayerArmourDamageEntry, 0, len(slots))
			for _, slot := range slots {
				if bitset&(1<<slot) != 0 {
					v := protocol.PlayerArmourDamageEntry{
						ArmourSlot: slot,
					}
					v2 := int32(0)
					io.Varint32(&v2)
					v.Damage = int16(v2)
					pk.List = append(pk.List, v)
				}
			}
			return
		}

		for _, entry := range pk.List {
			if entry.ArmourSlot == packet.PlayerArmourDamageFlagBody && !proto.IsProtoGTE(io, proto.ID712) {
				continue
			}
			bitset |= 1 << entry.ArmourSlot
		}
		io.Uint8(&bitset)
		for _, slot := range slots {
			if bitset&(1<<slot) != 0 {
				damage := int32(0)
				for _, entry := range pk.List {
					if entry.ArmourSlot == slot {
						damage = int32(entry.Damage)
						break
					}
				}
				io.Varint32(&damage)
			}
		}
		return
	}
	if proto.IsProtoLT(io, proto.ID859) {
		count := int32(len(pk.List))
		io.Varint32(&count)
		protocol.FuncIOSliceOfLen(io, uint32(count), &pk.List, func(io protocol.IO, entry *protocol.PlayerArmourDamageEntry) {
			slot := uint8(entry.ArmourSlot)
			io.Uint8(&slot)
			entry.ArmourSlot = int32(slot)
			io.Int16(&entry.Damage)
		})
		return
	}
	protocol.Slice(io, &pk.List)
}
