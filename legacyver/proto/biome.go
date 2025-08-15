package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// BiomeDefinition represents a biome definition in the game. This can be a vanilla biome or a completely
// custom biome.
type BiomeDefinition struct {
	// NameIndex represents the index of the biome name in the string list.
	NameIndex int16
	// BiomeID is the biome ID. This is optional and can be empty.
	BiomeID int16
	// Temperature is the temperature of the biome, used for weather, biome behaviours and sky colour.
	Temperature float32
	// Downfall is the amount that precipitation affects colours and block changes.
	Downfall float32
	// RedSporeDensity is the density of red spore precipitation visuals.
	RedSporeDensity float32
	// BlueSporeDensity is the density of blue spore precipitation visuals.
	BlueSporeDensity float32
	// AshDensity is the density of ash precipitation visuals.
	AshDensity float32
	// WhiteAshDensity is the density of white ash precipitation visuals.
	WhiteAshDensity float32
	// Depth ...
	Depth float32
	// Scale ...
	Scale float32
	// MapWaterColour is an ARGB value for the water colour on maps in the biome.
	MapWaterColour int32
	// Rain is true if the biome has rain, false if it is a dry biome.
	Rain bool
	// Tags are a list of indices of tags in the string list. These are used to group biomes together for
	// biome generation and other purposes.
	Tags protocol.Optional[[]uint16]
	// ChunkGeneration is optional information to assist in client-side chunk generation. Almost all servers
	// can and should leave this empty to greatly reduce the size of this packet. Only BDS and servers which
	// *exactly* match the vanilla chunk generation can benefit from this.
	ChunkGeneration protocol.Optional[protocol.BiomeChunkGeneration]
}

func (x *BiomeDefinition) Marshal(r protocol.IO) {
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
	r.Float32(&x.RedSporeDensity)
	r.Float32(&x.BlueSporeDensity)
	r.Float32(&x.AshDensity)
	r.Float32(&x.WhiteAshDensity)
	r.Float32(&x.Depth)
	r.Float32(&x.Scale)
	r.Int32(&x.MapWaterColour)
	r.Bool(&x.Rain)
	protocol.OptionalFunc(r, &x.Tags, func(s *[]uint16) {
		protocol.FuncSlice(r, s, r.Uint16)
	})
	protocol.OptionalMarshaler(r, &x.ChunkGeneration)
}

func (x *BiomeDefinition) FromLatest(bd protocol.BiomeDefinition) BiomeDefinition {
	x.NameIndex = bd.NameIndex
	x.BiomeID = bd.BiomeID
	x.Temperature = bd.Temperature
	x.Downfall = bd.Downfall
	x.RedSporeDensity = bd.RedSporeDensity
	x.BlueSporeDensity = bd.BlueSporeDensity
	x.AshDensity = bd.AshDensity
	x.WhiteAshDensity = bd.WhiteAshDensity
	x.Depth = bd.Depth
	x.Scale = bd.Scale
	x.MapWaterColour = bd.MapWaterColour
	x.Rain = bd.Rain
	x.Tags = bd.Tags
	x.ChunkGeneration = bd.ChunkGeneration
	return *x
}

func (x *BiomeDefinition) ToLatest() protocol.BiomeDefinition {
	return protocol.BiomeDefinition{
		NameIndex:        x.NameIndex,
		BiomeID:          x.BiomeID,
		Temperature:      x.Temperature,
		Downfall:         x.Downfall,
		RedSporeDensity:  x.RedSporeDensity,
		BlueSporeDensity: x.BlueSporeDensity,
		AshDensity:       x.AshDensity,
		WhiteAshDensity:  x.WhiteAshDensity,
		Depth:            x.Depth,
		Scale:            x.Scale,
		MapWaterColour:   x.MapWaterColour,
		Rain:             x.Rain,
		Tags:             x.Tags,
		ChunkGeneration:  x.ChunkGeneration,
	}
}
