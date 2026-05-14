package proto

import (
	_ "unsafe"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	RecipeFurnace     = 2
	RecipeFurnaceData = 3
)

func MarshalRecipe(w *protocol.Writer, recipe protocol.Recipe) {
	switch rec := recipe.(type) {
	case *protocol.ShapedRecipe:
		MarshalShapedRecipe(w, rec)
	case *protocol.ShapelessRecipe:
		MarshalShapelessRecipe(w, rec)
	case *protocol.ShulkerBoxRecipe:
		MarshalShapelessRecipe(w, &rec.ShapelessRecipe)
	case *protocol.ShapelessChemistryRecipe:
		MarshalShapelessRecipe(w, &rec.ShapelessRecipe)
	case *protocol.ShapedChemistryRecipe:
		MarshalShapedRecipe(w, &rec.ShapedRecipe)
	default:
		recipe.Marshal(w)
	}
}

func UnmarshalRecipe(r *protocol.Reader, recipe protocol.Recipe) {
	switch rec := recipe.(type) {
	case *protocol.ShapedRecipe:
		MarshalShapedRecipe(r, rec)
	case *protocol.ShapelessRecipe:
		MarshalShapelessRecipe(r, rec)
	case *protocol.ShulkerBoxRecipe:
		MarshalShapelessRecipe(r, &rec.ShapelessRecipe)
	case *protocol.ShapelessChemistryRecipe:
		MarshalShapelessRecipe(r, &rec.ShapelessRecipe)
	case *protocol.ShapedChemistryRecipe:
		MarshalShapedRecipe(r, &rec.ShapedRecipe)
	default:
		recipe.Unmarshal(r)
	}
}

func MarshalShapedRecipe(r protocol.IO, recipe *protocol.ShapedRecipe) {
	r.String(&recipe.RecipeID)
	r.Varint32(&recipe.Width)
	r.Varint32(&recipe.Height)
	protocol.FuncSliceOfLen(r, uint32(recipe.Width*recipe.Height), &recipe.Input, r.ItemDescriptorCount)
	protocol.FuncSlice(r, &recipe.Output, r.Item)
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	if IsProtoGTE(r, ID671) {
		r.Bool(&recipe.AssumeSymmetry)
	}
	if IsProtoGTE(r, ID685) {
		protocol.Single(r, &recipe.UnlockRequirement)
	}
	r.Varuint32(&recipe.RecipeNetworkID)
}

func MarshalShapelessRecipe(r protocol.IO, recipe *protocol.ShapelessRecipe) {
	r.String(&recipe.RecipeID)
	protocol.FuncSlice(r, &recipe.Input, r.ItemDescriptorCount)
	protocol.FuncSlice(r, &recipe.Output, r.Item)
	r.UUID(&recipe.UUID)
	r.String(&recipe.Block)
	r.Varint32(&recipe.Priority)
	if IsProtoGTE(r, ID685) {
		protocol.Single(r, &recipe.UnlockRequirement)
	}
	r.Varuint32(&recipe.RecipeNetworkID)
}

// FurnaceRecipe is a recipe that is specifically used for all kinds of furnaces. These recipes don't just
// apply to furnaces, but also blast furnaces and smokers.
type FurnaceRecipe struct {
	// InputType is the item type of the input item. The metadata value of the item is not used in the
	// FurnaceRecipe. Use FurnaceDataRecipe to allow an item with only one metadata value.
	InputType protocol.ItemType
	// Output is the item that is created as a result of smelting/cooking an item in the furnace.
	Output protocol.ItemStack
	// Block is the block name that is required to create the output of the recipe. The block is not prefixed
	// with 'minecraft:', so it will look like 'furnace' as an example.
	Block string
}

// FurnaceDataRecipe is a recipe specifically used for furnace-type crafting stations. It is equal to
// FurnaceRecipe, except it has an input item with a specific metadata value, instead of any metadata value.
type FurnaceDataRecipe struct {
	FurnaceRecipe
}

// Marshal ...
func (recipe *FurnaceRecipe) Marshal(w *protocol.Writer) {
	w.Varint32(&recipe.InputType.NetworkID)
	w.Item(&recipe.Output)
	w.String(&recipe.Block)
}

// Unmarshal ...
func (recipe *FurnaceRecipe) Unmarshal(r *protocol.Reader) {
	r.Varint32(&recipe.InputType.NetworkID)
	r.Item(&recipe.Output)
	r.String(&recipe.Block)
}

// Marshal ...
func (recipe *FurnaceDataRecipe) Marshal(w *protocol.Writer) {
	w.Varint32(&recipe.InputType.NetworkID)
	aux := int32(recipe.InputType.MetadataValue)
	w.Varint32(&aux)
	w.Item(&recipe.Output)
	w.String(&recipe.Block)
}

// Unmarshal ...
func (recipe *FurnaceDataRecipe) Unmarshal(r *protocol.Reader) {
	var dataValue int32
	r.Varint32(&recipe.InputType.NetworkID)
	r.Varint32(&dataValue)
	recipe.InputType.MetadataValue = uint32(dataValue)
	r.Item(&recipe.Output)
	r.String(&recipe.Block)
}

func lookupRecipe(recipeType int32, x *protocol.Recipe) bool {
	switch recipeType {
	case RecipeFurnace:
		*x = &FurnaceRecipe{}
		return true
	case RecipeFurnaceData:
		*x = &FurnaceDataRecipe{}
		return true
	default:
		return protoLookupRecipe(recipeType, x)
	}
}

func lookupRecipeType(x protocol.Recipe, recipeType *int32) bool {
	switch x.(type) {
	case *FurnaceRecipe:
		*recipeType = RecipeFurnace
		return true
	case *FurnaceDataRecipe:
		*recipeType = RecipeFurnaceData
		return true
	default:
		return protoLookupRecipeType(x, recipeType)
	}
}

//go:linkname protoLookupRecipe github.com/sandertv/gophertunnel/minecraft/protocol.lookupRecipe
func protoLookupRecipe(recipeType int32, x *protocol.Recipe) bool

//go:linkname protoLookupRecipeType github.com/sandertv/gophertunnel/minecraft/protocol.lookupRecipeType
func protoLookupRecipeType(x protocol.Recipe, recipeType *int32) bool
