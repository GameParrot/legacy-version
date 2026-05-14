package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func MarshalBiomeDefinition(r protocol.IO, x *protocol.BiomeDefinition) {
	r.Int16(&x.NameIndex)
	if IsProtoLT(r, ID827) {
		var opt protocol.Optional[int16]
		if x.BiomeID != -1 {
			opt = protocol.Option(x.BiomeID)
		}
		protocol.OptionalFunc(r, &opt, r.Int16)
		x.BiomeID, _ = opt.Value()
	} else {
		r.Int16(&x.BiomeID)
	}
	r.Float32(&x.Temperature)
	r.Float32(&x.Downfall)
	if IsProtoLT(r, ID844) {
		var f float32
		r.Float32(&f)
		r.Float32(&f)
		r.Float32(&f)
		r.Float32(&f)
		//r.Float32(&x.RedSporeDensity)
		//r.Float32(&x.BlueSporeDensity)
		//r.Float32(&x.AshDensity)
		//r.Float32(&x.WhiteAshDensity)
	}
	if IsProtoGTE(r, ID844) {
		r.Float32(&x.FoliageSnow)
	}
	r.Float32(&x.Depth)
	r.Float32(&x.Scale)
	r.Int32(&x.MapWaterColour)
	r.Bool(&x.Rain)
	protocol.OptionalFunc(r, &x.Tags, func(s *[]uint16) {
		protocol.FuncSlice(r, s, r.Uint16)
	})
	protocol.OptionalMarshaler(r, &x.ChunkGeneration)
}

func MarshalBiomeChunkGeneration(r protocol.IO, x *protocol.BiomeChunkGeneration) {
	protocol.OptionalMarshaler(r, &x.Climate)
	protocol.OptionalFunc(r, &x.ConsolidatedFeatures, func(s *[]protocol.BiomeConsolidatedFeature) {
		protocol.Slice(r, s)
	})
	protocol.OptionalMarshaler(r, &x.MountainParameters)
	protocol.OptionalFunc(r, &x.SurfaceMaterialAdjustments, func(s *[]protocol.BiomeElementData) {
		protocol.Slice(r, s)
	})
	protocol.OptionalMarshaler(r, &x.SurfaceMaterials)
	r.Bool(&x.HasDefaultOverworldSurface)
	r.Bool(&x.HasSwampSurface)
	r.Bool(&x.HasFrozenOceanSurface)
	r.Bool(&x.HasEndSurface)
	protocol.OptionalMarshaler(r, &x.MesaSurface)
	protocol.OptionalMarshaler(r, &x.CappedSurface)
	protocol.OptionalMarshaler(r, &x.OverworldRules)
	protocol.OptionalMarshaler(r, &x.MultiNoiseRules)
	protocol.OptionalFunc(r, &x.LegacyRules, func(s *[]protocol.BiomeConditionalTransformation) {
		protocol.Slice(r, s)
	})
	if IsProtoGTE(r, ID859) {
		protocol.OptionalFunc(r, &x.ReplacementsData, func(s *[]protocol.BiomeReplacementData) {
			protocol.Slice(r, s)
		})
	}
	if IsProtoGTE(r, ID924) {
		protocol.OptionalFunc(r, &x.VillageType, r.Uint8)
	}
}
