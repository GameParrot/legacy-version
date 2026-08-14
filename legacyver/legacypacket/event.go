package legacypacket

import (
	"fmt"

	"github.com/akmalfairuz/legacy-version/internal/typeconf"
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func Event(io protocol.IO, pk *packet.Event) {
	io.Varint64(&pk.EntityRuntimeID)
	io.EventType(&pk.Event)
	if proto.IsProtoGTE(io, proto.ID898) {
		io.Bool(&pk.UsePlayerID)
	} else {
		v := typeconf.BoolToByte(pk.UsePlayerID)
		io.Uint8(&v)
		pk.UsePlayerID = typeconf.ByteToBool(v)
	}
	if proto.IsProtoGTE(io, proto.ID898) {
		io.EventOrdinal(&pk.Event)
	}
	marshalEvent(io, pk.Event)
}

// marshalEvent serialises event data explicitly. Calling Event.Marshal here would
// bypass the versioned protocol layer whenever gophertunnel changes an event.
func marshalEvent(io protocol.IO, event protocol.Event) {
	switch x := event.(type) {
	case *protocol.AchievementAwardedEvent:
		io.Uint8(&x.AchievementID)
	case *protocol.EntityInteractEvent:
		io.ActorUniqueID(&x.InteractedEntityID)
		io.Uint8(&x.InteractionType)
		io.Varint32(&x.InteractionEntityType)
		io.Varint32(&x.EntityVariant)
		io.Uint8(&x.EntityColour)
	case *protocol.PortalBuiltEvent:
		io.Varint32(&x.DimensionID)
	case *protocol.PortalUsedEvent:
		io.Varint32(&x.FromDimensionID)
		io.Varint32(&x.ToDimensionID)
	case *protocol.MobKilledEvent:
		io.ActorUniqueID(&x.KillerEntityUniqueID)
		io.ActorUniqueID(&x.VictimEntityUniqueID)
		io.Varint32(&x.KillerEntityType)
		io.Varint32(&x.EntityDamageCause)
		io.Varint32(&x.VillagerTradeTier)
		io.String(&x.VillagerDisplayName)
	case *protocol.CauldronUsedEvent:
		io.Varuint32(&x.Colour)
		io.Int16(&x.PotionID)
		io.Int16(&x.FillLevel)
	case *protocol.PlayerDiedEvent:
		io.Varint32(&x.AttackerEntityID)
		io.Varint32(&x.AttackerVariant)
		io.Varint32(&x.EntityDamageCause)
		io.Bool(&x.InRaid)
	case *protocol.BossKilledEvent:
		io.ActorUniqueID(&x.BossEntityUniqueID)
		io.Varint32(&x.PlayerPartySize)
		io.Varint32(&x.InteractionEntityType)
	case *protocol.AgentCommandEvent:
		io.Varint32(&x.AgentResult)
		io.Varint32(&x.DataValue)
		io.String(&x.Command)
		io.String(&x.DataKey)
		io.String(&x.Output)
	case *protocol.SlashCommandExecutedEvent:
		io.Varint32(&x.SuccessCount)
		io.Varint32(&x.MessageCount)
		io.String(&x.CommandName)
		io.String(&x.OutputMessages)
	case *protocol.MobBornEvent:
		io.Varint32(&x.EntityType)
		io.Varint32(&x.Variant)
		io.Uint8(&x.Colour)
	case *protocol.CauldronInteractEvent:
		io.Uint8(&x.BlockInteractionType)
		io.Int16(&x.ItemID)
	case *protocol.ComposterInteractEvent:
		io.Uint8(&x.BlockInteractionType)
		io.Int16(&x.ItemID)
	case *protocol.BellUsedEvent:
		io.Int16(&x.ItemID)
	case *protocol.EntityDefinitionTriggerEvent:
		io.String(&x.EventName)
	case *protocol.RaidUpdateEvent:
		io.Varint32(&x.CurrentRaidWave)
		io.Varint32(&x.TotalRaidWaves)
		io.Bool(&x.WonRaid)
	case *protocol.TargetBlockHitEvent:
		io.Varint32(&x.RedstoneLevel)
	case *protocol.PiglinBarterEvent:
		io.Varint32(&x.ItemID)
		io.Bool(&x.WasTargetingBarteringPlayer)
	case *protocol.WaxedOrUnwaxedCopperEvent:
		io.Varint32(&x.CopperBlockID)
	case *protocol.CodeBuilderRuntimeActionEvent:
		io.String(&x.Action)
	case *protocol.CodeBuilderScoreboardEvent:
		io.String(&x.ObjectiveName)
		io.Varint32(&x.Score)
	case *protocol.ItemUsedEvent:
		io.Int16(&x.ItemID)
		io.Int32(&x.ItemAux)
		io.Int32(&x.UseMethod)
		io.Int32(&x.UseCount)
	case *protocol.AgentCreatedEvent, *protocol.PatternRemovedEvent, *protocol.FishBucketedEvent,
		*protocol.PetDiedEvent, *protocol.MovementAnomalyEvent, *protocol.MovementCorrectedEvent,
		*protocol.ExtractHoneyEvent, *protocol.StriderRiddenInLavaInOverworldEvent,
		*protocol.SneakCloseToSculkSensorEvent, *protocol.CarefulRestorationEvent:
		// These event variants have no payload.
	default:
		io.UnknownEnumOption(fmt.Sprintf("%T", event), "event type")
	}
}
