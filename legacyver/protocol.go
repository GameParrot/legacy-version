package legacyver

import (
	"github.com/akmalfairuz/legacy-version/legacyver/legacypacket"
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

var (
	packetPoolClient packet.Pool
	packetPoolServer packet.Pool

	packets map[uint32]func(io protocol.IO, pk packet.Packet)
)

func registerPacket[T packet.Packet](pk T, marshalFn func(io protocol.IO, pk T)) {
	packets[pk.ID()] = func(io protocol.IO, pk packet.Packet) {
		marshalFn(io, pk.(T))
	}
}

func registerPackets() {
	registerPacket(&packet.ActorEvent{}, legacypacket.ActorEvent)
	registerPacket(&packet.AddActor{}, legacypacket.AddActor)
	registerPacket(&packet.AddPlayer{}, legacypacket.AddPlayer)
	registerPacket(&packet.AddVolumeEntity{}, legacypacket.AddVolumeEntity)
	registerPacket(&packet.Animate{}, legacypacket.Animate)
	registerPacket(&packet.AnvilDamage{}, legacypacket.AnvilDamage)
	registerPacket(&packet.AvailableCommands{}, legacypacket.AvailableCommands)
	registerPacket(&packet.BiomeDefinitionList{}, legacypacket.BiomeDefinitionList)
	registerPacket(&packet.BlockActorData{}, legacypacket.BlockActorData)
	registerPacket(&packet.BlockEvent{}, legacypacket.BlockEvent)
	registerPacket(&packet.BookEdit{}, legacypacket.BookEdit)
	registerPacket(&packet.BossEvent{}, legacypacket.BossEvent)
	registerPacket(&packet.CameraAimAssistPresets{}, legacypacket.CameraAimAssistPresets)
	registerPacket(&packet.CameraAimAssist{}, legacypacket.CameraAimAssist)
	registerPacket(&packet.CameraInstruction{}, legacypacket.CameraInstruction)
	registerPacket(&packet.CameraPresets{}, legacypacket.CameraPresets)
	registerPacket(&packet.ChangeDimension{}, legacypacket.ChangeDimension)
	registerPacket(&packet.ClientBoundMapItemData{}, legacypacket.ClientBoundMapItemData)
	registerPacket(&packet.ClientCheatAbility{}, legacypacket.ClientCheatAbility)
	registerPacket(&packet.ClientMovementPredictionSync{}, legacypacket.ClientMovementPredictionSync)
	registerPacket(&packet.CodeBuilderSource{}, legacypacket.CodeBuilderSource)
	registerPacket(&packet.CommandBlockUpdate{}, legacypacket.CommandBlockUpdate)
	registerPacket(&packet.CommandOutput{}, legacypacket.CommandOutput)
	registerPacket(&packet.CommandRequest{}, legacypacket.CommandRequest)
	registerPacket(&packet.ContainerClose{}, legacypacket.ContainerClose)
	registerPacket(&packet.ContainerOpen{}, legacypacket.ContainerOpen)
	registerPacket(&packet.ContainerRegistryCleanup{}, legacypacket.ContainerRegistryCleanup)
	registerPacket(&packet.CorrectPlayerMovePrediction{}, legacypacket.CorrectPlayerMovePrediction)
	registerPacket(&packet.CraftingData{}, legacypacket.CraftingData)
	registerPacket(&packet.CreativeContent{}, legacypacket.CreativeContent)
	registerPacket(&packet.Disconnect{}, legacypacket.Disconnect)
	registerPacket(&packet.EditorNetwork{}, legacypacket.EditorNetwork)
	registerPacket(&packet.Emote{}, legacypacket.Emote)
	registerPacket(&packet.Event{}, legacypacket.Event)
	registerPacket(&packet.GameRulesChanged{}, legacypacket.GameRulesChanged)
	registerPacket(&packet.GraphicsOverrideParameter{}, legacypacket.GraphicsOverrideParameter)
	registerPacket(&packet.Interact{}, legacypacket.Interact)
	registerPacket(&packet.InventoryContent{}, legacypacket.InventoryContent)
	registerPacket(&packet.InventorySlot{}, legacypacket.InventorySlot)
	registerPacket(&packet.InventoryTransaction{}, legacypacket.InventoryTransaction)
	registerPacket(&packet.ItemRegistry{}, legacypacket.ItemRegistry)
	registerPacket(&packet.ItemStackRequest{}, legacypacket.ItemStackRequest)
	registerPacket(&packet.ItemStackResponse{}, legacypacket.ItemStackResponse)
	registerPacket(&packet.LecternUpdate{}, legacypacket.LecternUpdate)
	registerPacket(&packet.LevelSoundEvent{}, legacypacket.LevelSoundEvent)
	registerPacket(&packet.MobArmourEquipment{}, legacypacket.MobArmourEquipment)
	registerPacket(&packet.MobEffect{}, legacypacket.MobEffect)
	registerPacket(&packet.MobEquipment{}, legacypacket.MobEquipment)
	registerPacket(&packet.OpenSign{}, legacypacket.OpenSign)
	registerPacket(&packet.PartyChanged{}, legacypacket.PartyChanged)
	registerPacket(&packet.PlaySound{}, legacypacket.PlaySound)
	registerPacket(&packet.PlayerAction{}, legacypacket.PlayerAction)
	registerPacket(&packet.PlayerArmourDamage{}, legacypacket.PlayerArmourDamage)
	registerPacket(&packet.PlayerAuthInput{}, legacypacket.PlayerAuthInput)
	registerPacket(&packet.PlayerEnchantOptions{}, legacypacket.PlayerEnchantOptions)
	registerPacket(&packet.PlayerList{}, legacypacket.PlayerList)
	registerPacket(&packet.ResourcePackStack{}, legacypacket.ResourcePackStack)
	registerPacket(&packet.ResourcePacksInfo{}, legacypacket.ResourcePacksInfo)
	registerPacket(&packet.ServerBoundDiagnostics{}, legacypacket.ServerBoundDiagnostics)
	registerPacket(&packet.SetActorLink{}, legacypacket.SetActorLink)
	registerPacket(&packet.SetActorMotion{}, legacypacket.SetActorMotion)
	registerPacket(&packet.SetHud{}, legacypacket.SetHud)
	registerPacket(&packet.SetSpawnPosition{}, legacypacket.SetSpawnPosition)
	registerPacket(&packet.SetTitle{}, legacypacket.SetTitle)
	registerPacket(&packet.StopSound{}, legacypacket.StopSound)
	registerPacket(&packet.StructureBlockUpdate{}, legacypacket.StructureBlockUpdate)
	registerPacket(&packet.StructureTemplateDataRequest{}, legacypacket.StructureTemplateDataRequest)
	registerPacket(&packet.SubChunk{}, legacypacket.SubChunk)
	registerPacket(&packet.Text{}, legacypacket.Text)
	registerPacket(&packet.Transfer{}, legacypacket.Transfer)
	registerPacket(&packet.UpdateAbilities{}, legacypacket.UpdateAbilities)
	registerPacket(&packet.UpdateAttributes{}, legacypacket.UpdateAttributes)
	registerPacket(&packet.UpdateBlockSynced{}, legacypacket.UpdateBlockSynced)
	registerPacket(&packet.UpdateBlock{}, legacypacket.UpdateBlock)
	registerPacket(&packet.UpdateClientOptions{}, legacypacket.UpdateClientOptions)
	registerPacket(&packet.UpdatePlayerGameType{}, legacypacket.UpdatePlayerGameType)
	registerPacket(&packet.UpdateSubChunkBlocks{}, legacypacket.UpdateSubChunkBlocks)
	registerPacket(&packet.VoxelShapes{}, legacypacket.VoxelShapes)
}

func init() {
	packets = map[uint32]func(io protocol.IO, pk packet.Packet){}

	registerPackets()

	packetPoolClient = packet.NewClientPool()
	packetPoolServer = packet.NewServerPool()

	for pkId, cur := range packetPoolClient {
		packetPoolClient[pkId] = convertPacketFunc(pkId, cur)
	}

	for pkId, cur := range packetPoolServer {
		packetPoolServer[pkId] = convertPacketFunc(pkId, cur)
	}

}

func convertPacketFunc(pid uint32, cur func() packet.Packet) func() packet.Packet {
	if pid == packet.IDStartGame {
		return func() packet.Packet {
			return &legacypacket.TranslatedStartGame{Pk: cur().(*packet.StartGame)}
		}
	}
	if marshalFn, ok := packets[pid]; ok {
		return func() packet.Packet {
			return &translatedPacket{pk: cur(), marshalFn: marshalFn}
		}
	}
	return cur
}

type Protocol struct {
	ver string
	id  int32
}

func (p *Protocol) Ver() string {
	return p.ver
}

func (p *Protocol) ID() int32 {
	return p.id
}

func (p *Protocol) Packets(listener bool) packet.Pool {
	if listener {
		return packetPoolClient
	}
	return packetPoolServer
}

func (p *Protocol) NewReader(r minecraft.ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return proto.NewReader(protocol.NewReader(r, shieldID, enableLimits), p.id)
}

func (p *Protocol) NewWriter(w minecraft.ByteWriter, shieldID int32) protocol.IO {
	return proto.NewWriter(protocol.NewWriter(w, shieldID), p.id)
}

func (p *Protocol) ConvertToLatest(pk packet.Packet, conn *minecraft.Conn) []packet.Packet {
	return p.upgradePackets([]packet.Packet{pk}, conn)
}

func (p *Protocol) ConvertFromLatest(pk packet.Packet, conn *minecraft.Conn) []packet.Packet {
	return p.downgradePackets([]packet.Packet{pk}, conn)
}

func (p *Protocol) downgradePackets(pks []packet.Packet, conn *minecraft.Conn) []packet.Packet {
	for pkIndex, pk := range pks {
		if pk.ID() == packet.IDUpdateClientInputLocks {
			return []packet.Packet{}
		} else if pk.ID() == packet.IDStartGame {
			sg := pk.(*packet.StartGame)
			sg.GameVersion = p.ver
			sg.BaseGameVersion = p.ver
			pks[pkIndex] = &legacypacket.TranslatedStartGame{Pk: sg, Items: conn.GameData().Items}

		} else if marshalFn, ok := packets[pk.ID()]; ok {
			pks[pkIndex] = &translatedPacket{pk: pk, marshalFn: marshalFn}
		}
	}

	return pks
}

func (p *Protocol) upgradePackets(pks []packet.Packet, _ *minecraft.Conn) []packet.Packet {
	for pkIndex, pk := range pks {
		if pk.ID() == packet.IDStartGame {
			legacyStartGame := pk.(*legacypacket.TranslatedStartGame)
			pks[pkIndex] = legacyStartGame.Pk
		} else if translated, ok := pk.(*translatedPacket); ok {
			pks[pkIndex] = translated.pk
		}
	}
	return pks
}
