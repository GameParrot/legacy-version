package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type TranslatedStartGame struct {
	Pk    *packet.StartGame
	Items []protocol.ItemEntry
}

// ID ...
func (pk *TranslatedStartGame) ID() uint32 {
	return packet.IDStartGame
}

func (pk *TranslatedStartGame) Marshal(io protocol.IO) {
	StartGame(io, pk.Pk, pk.Items)
}

func StartGame(io protocol.IO, pk *packet.StartGame, items []protocol.ItemEntry) {
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.PlayerGameMode)
	io.Vec3(&pk.PlayerPosition)
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Int64(&pk.WorldSeed)
	io.Int16(&pk.SpawnBiomeType)
	io.String(&pk.UserDefinedBiomeName)
	io.Varint32(&pk.Dimension)
	io.Varint32(&pk.Generator)
	io.Varint32(&pk.WorldGameMode)
	if proto.IsProtoGTE(io, proto.ID671) {
		io.Bool(&pk.Hardcore)
	}
	io.Varint32(&pk.Difficulty)
	proto.IOUBlockPos(io, &pk.WorldSpawn)
	io.Bool(&pk.AchievementsDisabled)
	io.Varint32(&pk.EditorWorldType)
	io.Bool(&pk.CreatedInEditor)
	io.Bool(&pk.ExportedFromEditor)
	io.Varint32(&pk.DayCycleLockTime)
	io.Varint32(&pk.EducationEditionOffer)
	io.Bool(&pk.EducationFeaturesEnabled)
	io.String(&pk.EducationProductID)
	io.Float32(&pk.RainLevel)
	io.Float32(&pk.LightningLevel)
	io.Bool(&pk.ConfirmedPlatformLockedContent)
	io.Bool(&pk.MultiPlayerGame)
	io.Bool(&pk.LANBroadcastEnabled)
	io.Varint32(&pk.XBLBroadcastMode)
	io.Varint32(&pk.PlatformBroadcastMode)
	io.Bool(&pk.CommandsEnabled)
	io.Bool(&pk.TexturePackRequired)
	protocol.FuncSlice(io, &pk.GameRules, io.GameRuleLegacy)
	protocol.SliceUint32Length(io, &pk.Experiments)
	io.Bool(&pk.ExperimentsPreviouslyToggled)
	io.Bool(&pk.BonusChestEnabled)
	io.Bool(&pk.StartWithMapEnabled)
	io.Varint32(&pk.PlayerPermissions)
	io.Int32(&pk.ServerChunkTickRadius)
	io.Bool(&pk.HasLockedBehaviourPack)
	io.Bool(&pk.HasLockedTexturePack)
	io.Bool(&pk.FromLockedWorldTemplate)
	io.Bool(&pk.MSAGamerTagsOnly)
	io.Bool(&pk.FromWorldTemplate)
	io.Bool(&pk.WorldTemplateSettingsLocked)
	io.Bool(&pk.OnlySpawnV1Villagers)
	io.Bool(&pk.PersonaDisabled)
	io.Bool(&pk.CustomSkinsDisabled)
	io.Bool(&pk.EmoteChatMuted)
	io.String(&pk.BaseGameVersion)
	io.Int32(&pk.LimitedWorldWidth)
	io.Int32(&pk.LimitedWorldDepth)
	io.Bool(&pk.NewNether)
	protocol.Single(io, &pk.EducationSharedResourceURI)
	protocol.OptionalFunc(io, &pk.ForceExperimentalGameplay, io.Bool)
	io.Uint8(&pk.ChatRestrictionLevel)
	io.Bool(&pk.DisablePlayerInteractions)
	if proto.IsProtoGTE(io, proto.ID685) && proto.IsProtoLT(io, proto.ID924) {
		io.String(&pk.ServerID)
		io.String(&pk.WorldID)
		io.String(&pk.ScenarioID)
		if proto.IsProtoGTE(io, proto.ID818) {
			io.String(&pk.OwnerID)
		}
	}
	io.String(&pk.LevelID)
	io.String(&pk.WorldName)
	io.String(&pk.TemplateContentIdentity)
	io.Bool(&pk.Trial)
	proto.PlayerMoveSettings(io, &pk.PlayerMovementSettings)
	io.Int64(&pk.Time)
	io.Varint32(&pk.EnchantmentSeed)
	protocol.Slice(io, &pk.Blocks)
	if proto.IsProtoLT(io, proto.ID776) {
		legacyItems := make([]proto.LegacyItemRegistryEntry, len(items))
		for i, item := range items {
			legacyItems[i] = (&proto.LegacyItemRegistryEntry{}).FromLatest(item)
		}
		protocol.Slice(io, &legacyItems)
	}
	io.String(&pk.MultiPlayerCorrelationID)
	io.Bool(&pk.ServerAuthoritativeInventory)
	io.String(&pk.GameVersion)
	io.NBT(&pk.PropertyData, nbt.NetworkLittleEndian)
	io.Uint64(&pk.ServerBlockStateChecksum)
	io.UUID(&pk.WorldTemplateID)
	io.Bool(&pk.ClientSideGeneration)
	io.Bool(&pk.UseBlockNetworkIDHashes)
	if proto.IsProtoGTE(io, proto.ID827) && proto.IsProtoLT(io, proto.ID898) {
		v := false
		io.Bool(&v)
	}
	io.Bool(&pk.ServerAuthoritativeSound)
	if proto.IsProtoGTE(io, proto.ID924) {
		protocol.OptionalFuncIO(io, &pk.ServerJoinInformation, proto.MarshalServerJoinInformation)
		io.String(&pk.ServerID)
		io.String(&pk.ScenarioID)
		io.String(&pk.WorldID)
		io.String(&pk.OwnerID)
	}
}
