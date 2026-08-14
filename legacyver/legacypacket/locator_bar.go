package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func LocatorBar(io protocol.IO, pk *packet.LocatorBar) {
	protocol.FuncIOSlice(io, &pk.Waypoints, func(io protocol.IO, waypoint *protocol.LocatorBarWaypoint) {
		io.UUID(&waypoint.GroupHandle)
		io.Uint32(&waypoint.Waypoint.UpdateFlag)
		protocol.OptionalFunc(io, &waypoint.Waypoint.Visible, io.Bool)
		protocol.OptionalMarshaler(io, &waypoint.Waypoint.WorldPosition)
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.OptionalFunc(io, &waypoint.Waypoint.TexturePath, io.String)
			protocol.OptionalFunc(io, &waypoint.Waypoint.IconSize, io.Vec2)
		} else {
			var legacyTextureID protocol.Optional[uint32]
			protocol.OptionalFunc(io, &legacyTextureID, io.Uint32)
		}
		protocol.OptionalFunc(io, &waypoint.Waypoint.Colour, io.Int32)
		protocol.OptionalFunc(io, &waypoint.Waypoint.ClientPositionAuthority, io.Bool)
		if proto.IsProtoGTE(io, proto.ID2168) {
			protocol.OptionalFunc(io, &waypoint.Waypoint.ActorUniqueID, io.ActorUniqueID)
		} else {
			protocol.OptionalFunc(io, &waypoint.Waypoint.ActorUniqueID, io.Varint64)
		}
		io.Uint8(&waypoint.Action)
	})
}
