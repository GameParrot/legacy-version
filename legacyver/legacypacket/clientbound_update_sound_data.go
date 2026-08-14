package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func ClientboundUpdateSoundData(io protocol.IO, pk *packet.ClientboundUpdateSoundData) {
	io.Uint64(&pk.ServerSoundHandle)
	if proto.IsProtoGTE(io, proto.ID2168) {
		protocol.OptionalFunc(io, &pk.Stop, func(x *protocol.SoundDataUpdate) { proto.MarshalSoundDataUpdate(io, x) })
		protocol.OptionalFunc(io, &pk.SetVolume, func(x *protocol.SoundDataUpdate) { proto.MarshalSoundDataUpdate(io, x) })
		protocol.OptionalFunc(io, &pk.SetPitch, func(x *protocol.SoundDataUpdate) { proto.MarshalSoundDataUpdate(io, x) })
		protocol.OptionalFunc(io, &pk.Fade, func(x *protocol.SoundDataUpdate) { proto.MarshalSoundDataUpdate(io, x) })
		protocol.OptionalFunc(io, &pk.SeekTo, func(x *protocol.SoundDataUpdate) { proto.MarshalSoundDataUpdate(io, x) })
		protocol.OptionalFunc(io, &pk.Pause, func(x *protocol.SoundDataUpdate) { proto.MarshalSoundDataUpdate(io, x) })
		protocol.OptionalFunc(io, &pk.Resume, func(x *protocol.SoundDataUpdate) { proto.MarshalSoundDataUpdate(io, x) })
		return
	}
	event := "Stop"
	io.String(&event)
	if proto.IsReader(io) {
		if event != "Stop" {
			io.UnknownEnumOption(event, "sound data event")
			return
		}
		pk.Stop = protocol.Option(protocol.SoundDataUpdate{Type: protocol.SoundDataUpdateStop})
	}
}
