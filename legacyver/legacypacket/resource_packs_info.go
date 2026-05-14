package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ResourcePacksInfo(io protocol.IO, pk *packet.ResourcePacksInfo) {
	io.Bool(&pk.TexturePackRequired)
	if proto.IsProtoGTE(io, proto.ID662) {
		io.Bool(&pk.HasAddons)
	}
	io.Bool(&pk.HasScripts)
	if proto.IsProtoGTE(io, proto.ID766) {
		if proto.IsProtoGTE(io, proto.ID818) {
			io.Bool(&pk.ForceDisableVibrantVisuals)
		}
		io.UUID(&pk.WorldTemplateUUID)
		io.String(&pk.WorldTemplateVersion)
	}
	if proto.IsProtoLT(io, proto.ID729) {
		forcingServerPacks := false
		io.Bool(&forcingServerPacks)
		bps := []proto.BehaviourPackInfo{}
		protocol.SliceUint16Length(io, &bps)
	}
	protocol.SliceUint16Length(io, &pk.TexturePacks)
	if proto.IsProtoLT(io, proto.ID748) {
		if proto.IsReader(io) {
			packURLs := make([]protocol.PackURL, 0)
			protocol.Slice(io, &packURLs)
			for _, url := range packURLs {
				for _, t := range pk.TexturePacks {
					if url.UUIDVersion == t.UUID.String()+"_"+t.Version {
						t.DownloadURL = url.URL
					}
				}
			}
		} else {
			packURLs := make([]protocol.PackURL, 0)
			for _, t := range pk.TexturePacks {
				if t.DownloadURL != "" {
					packURLs = append(packURLs, protocol.PackURL{
						UUIDVersion: t.UUID.String() + "_" + t.Version,
						URL:         t.DownloadURL,
					})
				}
			}
			protocol.Slice(io, &packURLs)
		}
	}
}
