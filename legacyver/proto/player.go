package proto

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// PlayerListEntry is an entry found in the PlayerList packet. It represents a single player using the UUID
// found in the entry, and contains several properties such as the skin.
type PlayerListEntry struct {
	// UUID is the UUID of the player as sent in the Login packet when the client joined the server. It must
	// match this UUID exactly for the correct XBOX Live icon to show up in the list.
	UUID uuid.UUID
	// EntityUniqueID is the unique entity ID of the player. This ID typically stays consistent during the
	// lifetime of a world, but servers often send the runtime ID for this.
	EntityUniqueID int64
	// Username is the username that is shown in the player list of the player that obtains a PlayerList
	// packet with this entry. It does not have to be the same as the actual username of the player.
	Username string
	// XUID is the XBOX Live user ID of the player, which will remain consistent as long as the player is
	// logged in with the XBOX Live account.
	XUID string
	// PlatformChatID is an identifier only set for particular platforms when chatting (presumably only for
	// Nintendo Switch). It is otherwise an empty string, and is used to decide which players are able to
	// chat with each other.
	PlatformChatID string
	// BuildPlatform is the platform of the player as sent by that player in the Login packet.
	BuildPlatform int32
	// Skin is the skin of the player that should be added to the player list. Once sent here, it will not
	// have to be sent again.
	Skin protocol.Skin
	// Teacher is a Minecraft: Education Edition field. It specifies if the player to be added to the player
	// list is a teacher.
	Teacher bool
	// Host specifies if the player that is added to the player list is the host of the game.
	Host bool
	// SubClient specifies if the player that is added to the player list is a sub-client of another player.
	SubClient bool
	// PlayerColour is the colour of the player that is shown in UI elements, currently only used for the
	// player locator bar.
	PlayerColour color.RGBA
}

// Marshal encodes/decodes a PlayerListEntry.
func (x *PlayerListEntry) Marshal(r protocol.IO) {
	r.UUID(&x.UUID)
	r.Varint64(&x.EntityUniqueID)
	r.String(&x.Username)
	r.String(&x.XUID)
	r.String(&x.PlatformChatID)
	r.Int32(&x.BuildPlatform)
	protocol.Single(r, &x.Skin)
	r.Bool(&x.Teacher)
	r.Bool(&x.Host)
	r.Bool(&x.SubClient)
	if IsProtoGTE(r, ID800) {
		r.ARGB(&x.PlayerColour)
	}
}

// FromLatest ...
func (x *PlayerListEntry) FromLatest(y protocol.PlayerListEntry) PlayerListEntry {
	x.UUID = y.UUID
	x.EntityUniqueID = y.EntityUniqueID
	x.Username = y.Username
	x.XUID = y.XUID
	x.PlatformChatID = y.PlatformChatID
	x.BuildPlatform = y.BuildPlatform
	x.Skin = y.Skin
	x.Teacher = y.Teacher
	x.Host = y.Host
	x.SubClient = y.SubClient
	x.PlayerColour = y.PlayerColour
	return *x
}

// ToLatest ...
func (x *PlayerListEntry) ToLatest() protocol.PlayerListEntry {
	return protocol.PlayerListEntry{
		UUID:           x.UUID,
		EntityUniqueID: x.EntityUniqueID,
		Username:       x.Username,
		XUID:           x.XUID,
		PlatformChatID: x.PlatformChatID,
		BuildPlatform:  x.BuildPlatform,
		Skin:           x.Skin,
		Teacher:        x.Teacher,
		Host:           x.Host,
		SubClient:      x.SubClient,
		PlayerColour:   x.PlayerColour,
	}
}

// PlayerListRemoveEntry encodes/decodes a PlayerListEntry for removal from the list.
func PlayerListRemoveEntry(r protocol.IO, x *PlayerListEntry) {
	r.UUID(&x.UUID)
}

// PlayerMovementSettings represents the different server authoritative movement settings. These control how
// the client will provide input to the server.
type PlayerMovementSettings struct {
	// MovementType ...
	MovementType int32
	// RewindHistorySize is the amount of history to keep at maximum.
	RewindHistorySize int32
	// ServerAuthoritativeBlockBreaking specifies if block breaking should be sent through
	// packet.PlayerAuthInput or not.
	ServerAuthoritativeBlockBreaking bool
}

func (p *PlayerMovementSettings) FromLatest(x protocol.PlayerMovementSettings) PlayerMovementSettings {
	p.MovementType = 2
	p.RewindHistorySize = x.RewindHistorySize
	p.ServerAuthoritativeBlockBreaking = x.ServerAuthoritativeBlockBreaking
	return *p
}

func (p *PlayerMovementSettings) ToLatest() protocol.PlayerMovementSettings {
	return protocol.PlayerMovementSettings{
		RewindHistorySize:                p.RewindHistorySize,
		ServerAuthoritativeBlockBreaking: p.ServerAuthoritativeBlockBreaking,
	}
}

// PlayerMoveSettings reads/writes PlayerMovementSettings x to/from IO r.
func PlayerMoveSettings(r protocol.IO, x *PlayerMovementSettings) {
	if IsProtoLT(r, ID818) {
		r.Varint32(&x.MovementType)
	}
	r.Varint32(&x.RewindHistorySize)
	r.Bool(&x.ServerAuthoritativeBlockBreaking)
}
